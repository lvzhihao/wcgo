package wordcensor

/*
#include "wordcensor/platform.h"
#include "wordcensor/table.h"
#include "wordcensor/mm.h"
#include "wordcensor/hash.h"
#include "wordcensor/mmtable.h"
#include "wordcensor/docs.h"
#include "wordcensor/check.h"

#cgo linux pkg-config: glib-2.0
#cgo linux CFLAGS: -I/usr/local/wordcensor/include -pthread
#cgo linux LDFLAGS: -L/usr/local/wordcensor/lib/wordcensor -lwordcensor -lm
*/
import "C"
import "errors"
//import "fmt"

type WcStatus struct {
    Size int
}

func Create(filename string, flag string) (*C.wcMM, error) {
    var table *C.wcTable
    var err error
    C.wordcensor_create_table(&table)
    C.wordcensor_open_deny_file(C.CString(filename), table)
    var MM *C.wcMM
    C.wordcensor_mm_create(&MM, C.CString(flag))
    if C.wordcensor_mmtable_create(MM, table) != C.WORDCENSOR_SUCCESS{
        err = errors.New("mmtable create error!")
    }
    return MM, err
}

func Fetch(flag string) (*C.wcMM, error) {
    var MM *C.wcMM
    var err error
    if C.wordcensor_mm_fetch(&MM, C.CString(flag)) != C.WORDCENSOR_SUCCESS{
        err = errors.New("mmtable fetch error!")
    }
    return MM, err
}

func Status(MM *C.wcMM)  *WcStatus {
    ret := new(WcStatus)
    ret.Size = (int)(C.wordcensor_mm_size(MM))
    return ret
}

func Check(MM *C.wcMM, data string) (map[int]string, error) {
    var mmtable *C.wcmmTable
    var err error
    ret := make(map[int]string)
    if C.wordcensor_mmtable_fetch(MM, &mmtable) == C.WORDCENSOR_SUCCESS {
        var list *C.wcList
        var out *C.char
        var out_len C.int
        Cdata := C.CString(data) 
        if num := C.wordcensor_mm_check(MM, mmtable, Cdata, (C.int)(C.strlen(Cdata)), &out, &out_len, &list); num > 0 {
            var lt *C.wcList
            var res *C.wcResult
            for i:=0; C.wordcensor_list_get_current(list, &lt) == C.WORDCENSOR_SUCCESS; i++ {
                res = (*C.wcResult)(lt.val)
                ret[i] = C.GoString(res.string)
                //fmt.Printf("-->禁词: %s\n", C.GoString(res.string))
                //fmt.Printf("-->起始: %d\n", res.start)
                //fmt.Printf("-->长度: %d\n\n", res.len)
                C.wordcensor_list_next_item(&list)
            }
        }
    }else{
        err = errors.New("fetch error")
    }
    return ret, err
}


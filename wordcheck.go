package wordcheck

/*
#include "wordcheck/platform.h"
#include "wordcheck/table.h"
#include "wordcheck/mm.h"
#include "wordcheck/hash.h"
#include "wordcheck/mmtable.h"
#include "wordcheck/docs.h"
#include "wordcheck/check.h"

#cgo pkg-config: glib-2.0
#cgo CFLAGS: -I/usr/local/wordcheck/include -pthread
#cgo LDFLAGS: -L/usr/local/wordcheck/lib/wordcheck -lwordcheck -lm
*/
import "C"
import "errors"

//import "fmt"

type WcStatus struct {
	Size int
}

type WcResult struct {
	Word  string       `json:"word"`
	Info  WcResultInfo `json:"info"`
	Start int64        `json:"start"`
	Len   int64        `json:"len"`
}

type WcResultInfo struct {
	Weight int64 `json:"weight"`
}

func Create(filename string, flag string, memo string, size int) (*C.wcMM, error) {
	var table *C.wcTable
	var err error
	C.wordcheck_create_table(&table)
	C.wordcheck_open_deny_file(C.CString(filename), table)
	var MM *C.wcMM
	var MMInfo *C.wcMMInfo
	C.wordcheck_mminfo_create(&MMInfo, C.CString(memo), C.uint(size))
	C.wordcheck_mm_create(&MM, MMInfo, C.CString(flag))
	if C.wordcheck_mmtable_create(MM, table) != C.WORDCHECK_SUCCESS {
		err = errors.New("mmtable create error!")
	}
	return MM, err
}

func Fetch(flag string) (*C.wcMM, error) {
	var MM *C.wcMM
	var err error
	if C.wordcheck_mm_fetch(&MM, C.CString(flag)) != C.WORDCHECK_SUCCESS {
		err = errors.New("mmtable fetch error!")
	}
	return MM, err
}

func Status(MM *C.wcMM) *WcStatus {
	ret := new(WcStatus)
	ret.Size = (int)(C.wordcheck_mm_size(MM))
	return ret
}

func Check(MM *C.wcMM, data string) ([]WcResult, error) {
	var mmtable *C.wcmmTable
	var err error
	ret := make([]WcResult, 0)
	if C.wordcheck_mmtable_fetch(MM, &mmtable) == C.WORDCHECK_SUCCESS {
		var list *C.wcList
		var out *C.char
		var out_len C.int
		Cdata := C.CString(data)
		if num := C.wordcheck_mm_check(MM, mmtable, Cdata, (C.int)(C.strlen(Cdata)), &out, &out_len, &list); num > 0 {
			var lt *C.wcList
			var res *C.wcResult
			var wcResult WcResult
			for i := 0; C.wordcheck_list_get_current(list, &lt) == C.WORDCHECK_SUCCESS; i++ {
				res = (*C.wcResult)(lt.val)
				wcResult.Word = C.GoString(res.string)
				wcResult.Start = int64(res.start)
				wcResult.Len = int64(res.len)
				wcResult.Info.Weight = int64(res.info.weight)
				ret = append(ret, wcResult)
				//fmt.Printf("-->禁词: %s\n", C.GoString(res.string))
				//fmt.Printf("-->起始: %d\n", res.start)
				//fmt.Printf("-->长度: %d\n\n", res.len)
				C.wordcheck_list_next_item(&list)
			}
		}
	} else {
		err = errors.New("fetch error")
	}
	return ret, err
}

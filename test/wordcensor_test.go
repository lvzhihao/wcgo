package wordcensor_test

import wc "git.ishopex.cn/base/wordcensor"
import "fmt"

func main(){
    MM, err := wc.Create("demo.txt", "demo2")
    if err != nil{
    }
    fmt.Printf("mmtable size: %d\n", wc.Status(MM).Size)
    Map, err1 :=  wc.Check(MM, "隐藏在共和国的敌人")
    if err1 != nil{
    }
    fmt.Printf("Maps:\n%v\n", Map)
}

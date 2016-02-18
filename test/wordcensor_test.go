package wordcensor_test

import "testing"
import wc "github.com/lvzhihao/wcgo"
//import "fmt"

func TestCreate(t *testing.T){
    _, err := wc.Create("demo.txt", "wcgo_test", "wcgo_test", 67108864)
    if err == nil{
        t.Log("pass.. -_-! ")
    }else{
        t.Error("fail..")
    }
}

func TestFetch(t *testing.T){
    mm, err := wc.Fetch("wcgo_test")
    status := wc.Status(mm)
    if err == nil{
        t.Log("status: ", status)
        t.Log("pass.. -_-! ")
    }else{
        t.Error("fail..")
    }
}

func TestCheck(t *testing.T){
    MM, _ := wc.Fetch("wcgo_test")
    res, err := wc.Check(MM, "隐藏在共和国的敌人")
    if err == nil{
        t.Log(res)
        t.Log("pass.. -_-! ")
    }else{
        t.Error("fail..")
    }
}

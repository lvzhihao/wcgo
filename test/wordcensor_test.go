package wordcensor_test

import "testing"
import wc "git.ishopex.cn/base/wcgo"
//import "fmt"

func TestCreate(t *testing.T){
    _, err := wc.Create("demo.txt", "test_demo")
    if err != nil{
        t.Log("pass.. -_-! ")
    }else{
        t.Log("fail..")
    }
}

func TestFetch(t *testing.T){
    _, err := wc.Fetch("test_demo")
    if err != nil{
        t.Log("pass.. -_-! ")
    }else{
        t.Log("fail..")
    }
}

func TestCheck(t *testing.T){
    MM, _ := wc.Fetch("test_demo")
    _, err := wc.Check(MM, "隐藏在共和国的敌人")
    if err != nil{
        t.Log("pass.. -_-! ")
    }else{
        t.Log("fail..")
    }
}

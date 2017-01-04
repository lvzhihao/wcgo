package wordcheck_test

import (
	"strings"
	"testing"

	wc "github.com/lvzhihao/wcgo"
)

//import "fmt"

func BenchmarkCheckData(b *testing.B) {
	obj := wc.New(wc.Config{})
	obj.Load("wcgo_test")
	for n := 0; n < b.N; n++ {
		obj.Check("隐藏在共和国的敌人")
	}
}

func TestCreate(t *testing.T) {
	_, err := wc.Create("demo.txt", "wcgo_test", "wcgo_test", 67108864)
	if err == nil {
		t.Log("pass.. -_-! ")
	} else {
		t.Error("fail..")
	}
}

func TestLoad(t *testing.T) {
	mm, err := wc.Load("wcgo_test")
	status := wc.Info(mm)
	if err == nil {
		t.Log("status: ", status)
		t.Log("pass.. -_-! ")
	} else {
		t.Error("fail..")
	}
}

func TestCheck(t *testing.T) {
	MM, _ := wc.Load("wcgo_test")
	s, res, err := wc.Check(MM, "隐藏在共和国的敌人")
	if err == nil {
		t.Log(s)
		t.Log(res)
		t.Log("pass.. -_-! ")
	} else {
		t.Error("fail..")
	}
}

func TestInstanceLoad(t *testing.T) {
	obj := wc.New(wc.Config{})
	err := obj.Load("wcgo_test")
	if err == nil {
		t.Log("status: ", obj.Info())
		t.Log("pass.. -_-! ")
	} else {
		t.Error("fail..")
	}
}

func TestInstanceCheck(t *testing.T) {
	obj := wc.New(wc.Config{})
	obj.Load("wcgo_test")
	s, res, err := obj.Check("隐藏在共和国的敌人")
	if err == nil {
		t.Log(s)
		t.Log(res)
		t.Log("pass.. -_-! ")
	} else {
		t.Error("fail..")
	}
}

func TestInstanceConfig(t *testing.T) {
	var ro string = "#"
	var rl int = 5
	obj := wc.New(wc.Config{
		ReplaceOp:  ro,
		ReplaceLen: rl,
	})
	obj.Load("wcgo_test")
	s, res, err := obj.Check("隐藏在共和国的敌人")
	if err == nil && s == strings.Repeat(ro, rl)+"的敌人" {
		t.Log(s)
		t.Log(res)
		t.Log("pass.. -_-! ")
		obj.SetConfig(wc.Config{
			ReplaceOp: "%",
		})
		t.Log(obj.Check("隐藏在共和国的敌人"))
	} else {
		t.Error("fail..")
	}
}

package cycle

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestConstructorWithFullname(t *testing.T) {
	var aho CycleProc
	aho.Action = yaruyo
	aho.Flg = true
	aho.Time = 3
	DoProcess(aho)
	aho.Flg = false
	sleepTime := time.Duration(5) * time.Millisecond
	time.Sleep(sleepTime)
	if exists("hoge") != true {
		t.Error("cant do process")
	} else {
		os.Remove("hoge")
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func yaruyo() {
	fv := reflect.ValueOf(yaruyo)
	fmt.Println(runtime.FuncForPC(fv.Pointer()).Name())
	content := []byte("hello world\n")
	ioutil.WriteFile("hoge", content, os.ModePerm)
}

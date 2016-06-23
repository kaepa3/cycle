package main

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
	aho.action = yaruyo
	aho.flg = true
	aho.time = 3
	doProcess(aho)
	aho.flg = false
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

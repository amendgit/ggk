package ggk

import (
	"fmt"
	"runtime"
	"path/filepath"
)

func toimpl() {
	var pc, file, line, ok = runtime.Caller(1)
	if !ok {
		return
	}
	var _, filename = filepath.Split(file)
	var funcname = runtime.FuncForPC(pc).Name()
	fmt.Printf("toimpl \t %-22v \t %-5v \t %-32v\n", filename, line, funcname)
}

func warn(format string, a ...interface{}) {
	var pc, file, line, ok = runtime.Caller(1)
	if !ok {
		return
	}
	var _, filename = filepath.Split(file)
	var funcname = runtime.FuncForPC(pc).Name()
	var msg = fmt.Sprintf(format, a)
	fmt.Println("warn %v %v %v %v", filename, line, funcname, msg)
}

func errorf(format string, a ...interface{}) error {
	var pc, file, line, ok = runtime.Caller(1)
	if !ok {
		return nil
	}
	var _, filename = filepath.Split(file)
	var funcname = runtime.FuncForPC(pc).Name()
	var msg = fmt.Sprintf(format, a)
	return fmt.Errorf("%v %v %v %v", filename, line, funcname, msg)
}

package ggk

import (
	"fmt"
	"runtime"
)

func toimpl() {
	var pc, _, line, ok = runtime.Caller(1)
	if !ok {
		return
	}
	var fn = runtime.FuncForPC(pc).Name()
	fmt.Println("toimpl", fn, line)
}

func warning(format string, a ...interface{}) {
	var pc, file, line, ok = runtime.Caller(1)
	if !ok {
		return
	}
	var fn = runtime.FuncForPC(pc).Name()
	var msg = fmt.Sprintf(format, a)
	fmt.Println("warning %v %v %v %v", file, fn, line, msg)
}

func errorf(format string, a ...interface{}) error {
	var pc, file, line, ok = runtime.Caller(1)
	if !ok {
		return nil
	}
	var fn = runtime.FuncForPC(pc).Name()
	var msg = fmt.Sprintf(format, a)
	return fmt.Errorf("%v %v %v %v", file, fn, line, msg)
}

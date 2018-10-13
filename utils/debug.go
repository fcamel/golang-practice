package utils

import (
	"fmt"
	"path"
	"runtime"
)

func Trace(format string, a ...interface{}) {
	function, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf("> %s:%d %s:", path.Base(file), line, runtime.FuncForPC(function).Name())
	msg := fmt.Sprintf(format, a...)
	fmt.Println(info, msg)
}

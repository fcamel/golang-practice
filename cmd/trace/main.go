package main

import (
	"github.com/fcamel/golang-practice/utils"
)

type myType struct {
}

func (t myType) hello() {
	utils.Trace("")
}

func foo() {
	utils.Trace("begin")
	defer utils.Trace("end")
	bar()
}

func bar() {
	utils.Trace("Hello %d", 101)
	var t myType
	t.hello()
}

func main() {
	foo()
}

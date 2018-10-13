package main

import (
	"github.com/fcamel/golang-practice/utils"
)

func foo() {
	utils.Trace("begin")
	defer utils.Trace("end")
	bar()
}

func bar() {
	utils.Trace("Hello %d", 101)
}

func main() {
	foo()
}

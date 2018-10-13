package utils_test

import (
	"utils"
)

func ExampleTrace() {
	utils.Trace("hello %d", 101)
	// Output to stdout with something like:
	// > example_test.go:8 XXX.ExampleTrace: hello 7
}

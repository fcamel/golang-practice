package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

var n = 1000000000
var ctx = context.Background()

func a() {
	b()
	c()
}

func b() {
	d()
	s := 0
	for i := 0; i < n; i++ {
		s += i
	}
	fmt.Println("b", s)
}

func c() {
	time.Sleep(3 * time.Second)
	d()
}

func d() {
	s := 0
	for i := 0; i < n*2; i++ {
		s += i
	}
	fmt.Println("d", s)
}

func main() {
	// CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// Trace
	f, err = os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	// Run program.
	a()
}

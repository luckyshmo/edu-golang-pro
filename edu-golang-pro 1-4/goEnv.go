package main

import (
	"fmt"
	"runtime"
)

func runtimeInfo() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

	//go env - all env for current machine
	// go tool compile -W nodeTree.go - compiled GO code
	//go build -x - info about build
}

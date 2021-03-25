package main

import (
	"fmt"
	"runtime"
	"time"
)

func notWorking() { //chanels have same behavior //TODO
	for i := 0; i <= 20; i++ {
		go func() {
			fmt.Print(i, " ") //even compiler know about this problem
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
}

func working() {
	for i := 0; i <= 20; i++ {
		i := i //lol, magic :D
		go func() {
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
}

//infinity loop func
func oneMoreStrangeBehavior() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ { //explanation pg 487-488
		}
	}()
	fmt.Println("Leaving goroutine!")
	runtime.Gosched()
	runtime.GC()

	fmt.Println("Good bye!")
}

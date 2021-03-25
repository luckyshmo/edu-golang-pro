package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select { //all cases executes at the same time
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		//this branch is like default. It's wait until 4 secnods pass and executes
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		}
	}
}

func selectTest() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	n := 1000
	fmt.Printf("Going to create %d random numbers.\n", n)
	go gen(0, 2*n, createNumber, end)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
	end <- true
}

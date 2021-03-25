package main

import (
	"fmt"
	"sync"
	"time"
)

func to1() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout c1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout c2")
	}
}

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func to2() {

	var w sync.WaitGroup
	w.Add(1)

	t := 6000

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
}

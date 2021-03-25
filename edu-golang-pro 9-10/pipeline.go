package main

import (
	"fmt"
	"math/rand"
	"time"
)

//then second and first try to read and write this global var they are in 'race' state
// to fix this we will use signal channel and 'select' key word

var DATA = make(map[int]bool)

var signal chan struct{}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		select { //select is kind of switch, but for channels
		case <-signal:
			close(out)
			return
		case out <- random(min, max):
		}
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		_, ok := DATA[x]
		if ok {
			signal <- struct{}{}
		} else {
			fmt.Print(x, " ")
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d.\n", sum)
}

func pipeline() {

	n1 := 1
	n2 := 1000
	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d.\n", n1, n2)
		return
	}

	signal = make(chan struct{})

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(B, A)
	third(B)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomTest() {
	MIN := 0
	MAX := 100
	TOTAL := 100
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		myrand := random(MIN, MAX)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}

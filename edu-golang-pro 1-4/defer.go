package main

import (
	"fmt"
	"log"
	"os"
)

// если сначала отложить функцию f1() ,
// затем функцию f2() и после нее — функцию f3() в одной и той же внешней функ-
// ции, то, когда эта внешняя функция будет возвращать значение, функция f3()
// выполнится первой, функция f2() — второй, а функция f1() — последней.

func one(aLog *log.Logger) {
	aLog.Println("-- FUNCTION one ------")
	defer aLog.Println("-- FUNCTION one ------")

	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

func two(aLog *log.Logger) {
	aLog.Println("---- FUNCTION two")
	defer aLog.Println("FUNCTION two ------")

	for i := 10; i > 0; i-- {
		aLog.Println(i)
	}
}

func deferLog() {
	f, err := os.OpenFile("./mGo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "logDefer ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")

	one(iLog)
	two(iLog)
}

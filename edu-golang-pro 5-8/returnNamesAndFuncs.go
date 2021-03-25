package main

import (
	"fmt"
	"os"
	"strconv"
)

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

func namedFunc() {
	// 	использование именованных переменных
	// очень полезно, если вам нужно вернуть значение из функции после
	// обработки паники
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}

	a1, _ := strconv.Atoi(arguments[1])
	a2, _ := strconv.Atoi(arguments[2])

	fmt.Println(minMax(a1, a2))
	min, max := minMax(a1, a2)
	fmt.Println(min, max)

	fmt.Println(namedMinMax(a1, a2))
	min, max = namedMinMax(a1, a2)
	fmt.Println(min, max)
}

func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func funcTest() {
	i := funReturnFun()
	j := funReturnFun()

	fmt.Println("1:", i())
	fmt.Println("2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
	fmt.Println("3:", i())
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

type person struct {
	Name  string //truly random chars 🗼, random size
	Sname []byte //random size
	Runes [3]int
	c     complex64 //! should be exported
	P     point3D
}

type point3D struct {
	X, Y, Z int8
	S       float32
}

var ran = rand.New(rand.NewSource(time.Now().Unix()))

func bi() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover!") //TODO how about fucking err????
		}
	}()

	myValues := reflect.TypeOf(point3D{})
	x, _ := quick.Value(myValues, ran)
	fmt.Println(x)

	custom()

	fmt.Println("Finished")
}

func custom() {
	kekP := reflect.TypeOf(person{})
	y, ok := quick.Value(kekP, ran)
	if !ok {
		log.Fatal(ok)
	}
	fmt.Println(y)
}

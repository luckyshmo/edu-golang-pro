package main

import (
	"fmt"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func reflection() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xType)
	// fmt.Printf("The type of x is %s.\n", reflect.TypeOf(x))

	A := a{100, 200.12, "Struct a"}
	// B := b{1, 2, "Struct b", -1.2}
	var r reflect.Value

	r = reflect.ValueOf(&A).Elem()
	// r = reflect.ValueOf(&B).Elem()

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are:\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and value %v\n", r.Field(i).Interface())
	}
}

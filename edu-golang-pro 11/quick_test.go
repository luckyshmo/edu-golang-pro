package main

import (
	"testing"
	"testing/quick"
)

var N = 100000

func Add(x, y uint16) uint16 {
	var i uint16
	for i = 0; i < x; i++ {
		y++
	}
	return y
}

//go test -v quick* //! cash
//go test -v quick* -count=1 //! no cash
//go test quick* -v -timeout 2s //! fist will pass, second will fail

func TestWithSystem(t *testing.T) {
	condition := func(a, b uint16) bool {
		return Add(a, b) == (b + a)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestWithItself(t *testing.T) {
	condition := func(a, b uint16) bool {
		return Add(a, b) == Add(b, a)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

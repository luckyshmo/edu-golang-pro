package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverseBT(t *Tree) {
	if t == nil {
		return
	}
	traverseBT(t.Left)
	fmt.Print(t.Value, " ")
	traverseBT(t.Right)
}

func createBT(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insertBT(t, temp)
	}
	return t
}

func insertBT(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insertBT(t.Left, v)
		return t
	}
	t.Right = insertBT(t.Right, v)
	return t
}

func mainBintree() { //TODO where to use binTree?
	tree := createBT(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverseBT(tree)
	fmt.Println()
	tree = insertBT(tree, -10)
	tree = insertBT(tree, -2)
	traverseBT(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}

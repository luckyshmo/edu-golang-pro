package main

import (
	"fmt"
)

const SIZE = 30

type NodeH struct {
	Value int
	Next  *NodeH
}

type HashTable struct {
	Table map[int]*NodeH
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insertTable(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := NodeH{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverseTable(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func mainHashTable() {
	table := make(map[int]*NodeH, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insertTable(hash, i)
	}
	traverseTable(hash)

	for i := 10; i < 125; i++ {
		if !lookup(hash, i) {
			fmt.Println(i, "is not in the hash table!")
		}
	}
}

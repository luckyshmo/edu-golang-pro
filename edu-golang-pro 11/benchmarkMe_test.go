package main

import (
	"testing"
)

// go test -benchmem -bench=. benchmarkMe_test.go

func fibo1b(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo1b(n-1) + fibo1b(n-2)
	}
}

func fibo2b(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo2b(n-1) + fibo2b(n-2)
}

func fibo3b(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i == 0 {
			f = 0
		} else if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

//Test:

var result int

func benchmarkfibo1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo1b(n)
	}
	result = r
}

func benchmarkfibo2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo2b(n)
	}
	result = r
}

func benchmarkfibo3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo3b(n)
	}
	result = r
}

func Benchmark30fibo1(b *testing.B) {
	benchmarkfibo1(b, 30)
}

func Benchmark30fibo2(b *testing.B) {
	benchmarkfibo2(b, 30)
}

func Benchmark30fibo3(b *testing.B) {
	benchmarkfibo3(b, 30)
}

func Benchmark50fibo1(b *testing.B) {
	benchmarkfibo1(b, 50)
}

func Benchmark50fibo2(b *testing.B) {
	benchmarkfibo2(b, 50)
}

func Benchmark50fibo3(b *testing.B) {
	benchmarkfibo3(b, 50)
}

// This is a correct benchmark function
func BenchmarkFiboIV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibo3b(10)
	}
}

// This is also a correct benchmark function
func BenchmarkFiboIII(b *testing.B) {
	_ = fibo3b(b.N)
}

// This benchmark function never converges
// func BenchmarkFiboI(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		_ = fibo1(i)
//	}
// }

// This benchmark function never converges
// func BenchmarkFiboII(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		_ = fibo2(b.N)
//	}
// }

package codeCover

//go test --cover

// go test -coverprofile=coverage.out
// go tool cover -func=coverage.out

func fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo1(n-1) + fibo1(n-2)
	}
}

func fibo2(n int) int {
	if n >= 0 { //mistake
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo1(n-1) + fibo1(n-2)
	}
}

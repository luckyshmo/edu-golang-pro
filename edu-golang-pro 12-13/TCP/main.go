package main

import (
	"fmt"
	"time"
)

func main() {
	// go TCPServer(":8001")
	// time.Sleep(2 * time.Second)
	// TCPClient(":8001", []string{"Hi server!", "He-hey"})

	// go TCPServer2("localhost:8001")
	// time.Sleep(time.Second * 2)
	// TCPClient2("localhost:8001")

	//testFibo()

	testKV()
	// go kvTCP(":9191")
	// time.Sleep(2 * time.Second)
	// TCPClient(":9191", []string{"PRINT", "STOP"})
}

func testKV() {
	go kvTCP(":9191")
	time.Sleep(2 * time.Second)
	t := time.Now()
	for j := 0; j < 25; j++ {
		for i := 0; i < 25; i++ {
			go TCPClient(":9191", []string{"ADD " + fmt.Sprint(j*i+i) + " kek, lol, wow"})
		}
		time.Sleep(time.Millisecond * 2000)
	}
	time.Sleep(time.Second)
	TCPClient(":9191", []string{"ADD " + fmt.Sprint(2000) + " kek, lol, wow"})
	TCPClient(":9191", []string{"STOP"})
	// TCPClient(":9191", []string{"100", "123", "123", "1000", "110000", "110000", "110000", "110000", "110000"})
	fmt.Println(time.Since(t))
}

func testFibo() {
	go fiboTCP(":9191")
	time.Sleep(2 * time.Second)
	t := time.Now()
	for i := 0; i < 64; i++ {
		go TCPClient(":9191", []string{"100", "123", "123", "1000", "110000", "110000", "110000", "110000", "110000", "11000000"})
	}
	TCPClient(":9191", []string{"100", "123", "123", "1000", "110000", "110000", "110000", "110000", "110000"})
	fmt.Println(time.Since(t))
}

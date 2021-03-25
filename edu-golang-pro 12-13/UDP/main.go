package main

import "time"

func main() {
	go UDPServer(":9090")
	time.Sleep(time.Second * 2)
	UDPClient("localhost:9090")
}

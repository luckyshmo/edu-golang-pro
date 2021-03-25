package main

import (
	"fmt"
	"time"
)

func main() {

	// httpsClient("https://www.google.com")
	// httpsServer()

	signal := make(chan bool)

	go tlsServer(signal)

	for {
		select {
		case <-signal:
			tlsClient("https://localhost:1443")
		case <-time.After(5 * time.Second):
			fmt.Println("quit")
			return
		}
	}

}

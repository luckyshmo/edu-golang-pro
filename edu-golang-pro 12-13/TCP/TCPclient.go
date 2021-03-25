package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func TCPClient(hostPort string, msg []string) {

	c, err := net.Dial("tcp", hostPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, m := range msg {
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print(">> ")
		// text, err := reader.ReadString('\n')
		// if err != nil {
		// 	log.Fatal(err)
		// }
		text := m
		fmt.Fprintf(c, text+"\n")

		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("->: " + message)
		// time.Sleep(time.Second)
		// if strings.TrimSpace(string(text)) == "STOP" {
		// 	fmt.Println("TCP client exiting...")
		// 	return
		// }
	}
}

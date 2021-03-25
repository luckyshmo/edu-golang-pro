package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func TCPClient2(CONNECT string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", CONNECT)
	if err != nil {
		fmt.Println("ResolveTCPAddr:", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP:", err.Error())
		return
	}

	for {
		text := "Hi Server!"
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			conn.Close()
			return
		}
		time.Sleep(time.Second)
	}
}

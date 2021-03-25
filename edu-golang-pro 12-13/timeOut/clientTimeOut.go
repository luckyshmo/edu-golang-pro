package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func ct(URL string) {
	t := http.Transport{
		Dial: Timeout,
	}

	client := http.Client{
		Transport: &t,
	}

	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

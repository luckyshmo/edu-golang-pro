package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//go run webClient.go https://google.com

func simpleClient(URL string) {
	data, err := http.Get(URL)

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

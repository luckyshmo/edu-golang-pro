package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func at(URL string) {

	client := http.Client{
		Timeout: timeout,
	}
	client.Get(URL)

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

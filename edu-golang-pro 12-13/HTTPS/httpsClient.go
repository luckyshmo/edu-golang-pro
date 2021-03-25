package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpsClient(URL string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))

	fmt.Println(s)
}

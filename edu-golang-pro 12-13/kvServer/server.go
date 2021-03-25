package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8001"

func main() {
	err := load()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/list", listAll)
	http.HandleFunc("/insert", insertElement)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}

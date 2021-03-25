package main

import (
	"fmt"
	"net/http"
)

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is an example HTTPS server!\n")
}

func httpsServer() {
	var PORT = ":1443"
	http.HandleFunc("/", Default)
	fmt.Println("Listening to port number", PORT)

	err := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("ListenAndServeTLS: ", err)
		return
	}
}

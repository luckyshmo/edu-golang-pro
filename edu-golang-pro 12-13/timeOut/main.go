package main

import "time"

var timeout = time.Duration(time.Second)

func main() {
	ct("http://localhost:8001")
	at("http://localhost:8001")
	st()
}

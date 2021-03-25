package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var countRequests int32

func handleAll(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&countRequests, 1)
	// fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	// fmt.Printf("Served: %s\n", r.Host)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	temp := atomic.LoadInt32(&countRequests)
	fmt.Println("Count:", temp)
	fmt.Fprintf(w, "<h1 align=\"center\">%d</h1>", countRequests)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

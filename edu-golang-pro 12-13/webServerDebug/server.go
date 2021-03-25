package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime"
)

const PORT = ":8001"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() - 4)

	r := http.NewServeMux()
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/getCounter", getCounter)
	r.HandleFunc("/", handleAll) //default handler

	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

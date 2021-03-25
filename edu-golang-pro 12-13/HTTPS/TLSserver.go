package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

var PORT = ":1443"

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world!\n"))
}

func tlsServer(signal chan<- bool) {
	caCert, err := ioutil.ReadFile("client.crt")
	if err != nil {
		fmt.Println(err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	srv := &http.Server{
		Addr:      PORT,
		Handler:   &handler{},
		TLSConfig: cfg,
	}
	fmt.Println("Listening to port number", PORT)
	signal <- true
	fmt.Println(srv.ListenAndServeTLS("server.crt", "server.key"))
}

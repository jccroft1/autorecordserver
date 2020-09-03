package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port        string
	publicCert  string
	privateCert string
)

func main() {
	flag.StringVar(&port, "port", ":443", "custom port for server to listen on")
	flag.StringVar(&publicCert, "public_cert", "/certs/cert.pem", "path for public certificate")
	flag.StringVar(&privateCert, "private_cert", "/certs/privkey.pem", "path for private certificate")
	flag.Parse()

	http.HandleFunc("/", hello)

	http.ListenAndServeTLS(port, publicCert, privateCert, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!\n")
}

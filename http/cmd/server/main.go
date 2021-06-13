package main

import (
	"net/http"

	"github.com/taogilaaa/trace-sandbox/http/pkg/saleorder"
)

func main() {
	httpServer := saleorder.NewHTTPServer()

	http.HandleFunc("/hello", httpServer.Hello)

	http.ListenAndServe(":50042", nil)
}

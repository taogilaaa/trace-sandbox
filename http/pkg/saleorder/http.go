package saleorder

import (
	"fmt"
	"net/http"
)

type httpServer struct {
}

// NewHTTPServer creates a http server request handler.
func NewHTTPServer() *httpServer {
	return &httpServer{}
}

func (hs *httpServer) Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

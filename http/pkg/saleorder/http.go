package saleorder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/taogilaaa/trace-sandbox/http/internal/proto/sandbox_sales_v1"
)

type httpServer struct {
	sosc sandbox_sales_v1.SaleOrderServiceClient
}

// NewHTTPServer creates a http server request handler.
func NewHTTPServer(soClient sandbox_sales_v1.SaleOrderServiceClient) *httpServer {
	return &httpServer{soClient}
}

func (hs *httpServer) Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func (hs *httpServer) SaleOrder(w http.ResponseWriter, req *http.Request) {
	saleOrderID := int32(0)
	if value, err := strconv.Atoi(path.Base(req.URL.Path)); err == nil {
		saleOrderID = int32(value)
	}

	response, err := hs.sosc.GetSaleOrder(context.Background(), &sandbox_sales_v1.GetSaleOrderRequest{Id: saleOrderID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonMessage, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, string(jsonMessage))
}

func (hs *httpServer) SaleOrders(w http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get("email")
	response, err := hs.sosc.GetSaleOrders(context.Background(), &sandbox_sales_v1.GetSaleOrdersRequest{Email: email})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonMessage, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, string(jsonMessage))
}

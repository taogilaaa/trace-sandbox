package saleorder

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/taogilaaa/trace-sandbox/http/internal/proto/sandbox_sales_v1"
)

type Messager interface {
	SendMessage(ctx context.Context, channel string, message interface{}) error
}
type httpServer struct {
	sosc sandbox_sales_v1.SaleOrderServiceClient
	m    Messager
}

// NewHTTPServer creates a http server request handler.
func NewHTTPServer(soClient sandbox_sales_v1.SaleOrderServiceClient, messager Messager) *httpServer {
	return &httpServer{soClient, messager}
}

func (hs *httpServer) Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func (hs *httpServer) SaleOrder(w http.ResponseWriter, req *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "http.SaleOrder")
	defer span.Finish()

	saleOrderID := int32(0)
	if value, err := strconv.Atoi(path.Base(req.URL.Path)); err == nil {
		saleOrderID = int32(value)
	}

	response, err := hs.sosc.GetSaleOrder(ctx, &sandbox_sales_v1.GetSaleOrderRequest{Id: saleOrderID})
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
	if req.Method == http.MethodPost {
		hs.CreateSaleOrders(w, req)
		return
	}

	hs.GetSaleOrders(w, req)
}

func (hs *httpServer) GetSaleOrders(w http.ResponseWriter, req *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "http.GetSaleOrders")
	defer span.Finish()

	email := req.URL.Query().Get("email")
	response, err := hs.sosc.GetSaleOrders(ctx, &sandbox_sales_v1.GetSaleOrdersRequest{Email: email})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonMessage, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, string(jsonMessage))
}

func (hs *httpServer) CreateSaleOrders(w http.ResponseWriter, req *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "http.CreateSaleOrders")
	defer span.Finish()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var newSaleOrder CreateSaleOrder
	err = json.Unmarshal(body, &newSaleOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = hs.m.SendMessage(ctx, NatsChannel, newSaleOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "order placed")
}

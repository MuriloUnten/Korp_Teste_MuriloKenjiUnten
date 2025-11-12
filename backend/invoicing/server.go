package invoicing

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MuriloUnten/Korp_Teste_MuriloKenjiUnten/internal"
)

const (
	PRODUCT_SERVICE_URL = "localhost:8081"
)

type APIServer struct {
	addr string
	svc Service
}
func NewAPIServer(addr, productUrl string) *APIServer {
	s := &APIServer{
		addr: addr,
		svc: NewInvoicingService(PRODUCT_SERVICE_URL),
	}

	http.HandleFunc("GET /invoices", internal.MakeHandler(s.handleGetInvoices))
	http.HandleFunc("GET /invoices/{id}", internal.MakeHandler(s.handleGetInvoiceByNumber))
	http.HandleFunc("POST /invoices", internal.MakeHandler(s.handleCreateInvoice))
	http.HandleFunc("PUT /invoices/{id}/close", internal.MakeHandler(s.handleCloseInvoice))

	return s
}

func (s *APIServer) Serve() {
	fmt.Println("Gateway Server running at: ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, nil))
}

func (s *APIServer) handleGetInvoices(w http.ResponseWriter, r *http.Request) error {
	invoices, err := s.svc.GetInvoices(context.TODO())
	if err != nil {

	}

	return internal.WriteJSON(w, http.StatusOK, invoices)
}

func (s *APIServer) handleGetInvoiceByNumber(w http.ResponseWriter, r *http.Request) error {
	id, err := internal.GetPathId("id", r)
	if err != nil {
		return internal.InvalidPathIdentifier()
	}

	resp, err := s.svc.GetInvoiceByNumber(context.TODO(), id)
	if err != nil {
		
	}

	return internal.WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleCreateInvoice(w http.ResponseWriter, r *http.Request) error {
	var req internal.CreateInvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return internal.InvalidRequestBody()
	}

	resp, err := s.svc.CreateInvoice(context.TODO(), req)
	if err != nil {
		
	}

	return internal.WriteJSON(w, http.StatusCreated, resp)
}

func (s *APIServer) handleCloseInvoice(w http.ResponseWriter, r *http.Request) error {
	id, err := internal.GetPathId("id", r)
	if err != nil {
		return internal.InvalidPathIdentifier()
	}

	invoice, err := s.svc.CloseInvoice(context.TODO(), id)
	if err != nil {

	}

	return internal.WriteJSON(w, http.StatusOK, invoice)
}

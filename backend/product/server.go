package product

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MuriloUnten/Korp_Teste_MuriloKenjiUnten/internal"
)

type APIServer struct {
	addr string
	svc Service
}

func NewAPIServer(addr string) *APIServer {
	s := &APIServer{
		addr: addr,
		svc: NewProductService(),
	}

	http.HandleFunc("GET /products", internal.MakeHandler(s.handleGetProducts))
	http.HandleFunc("GET /products/{id}", internal.MakeHandler(s.handleGetProductById))
	http.HandleFunc("POST /products", internal.MakeHandler(s.handleCreateProduct))

	return s
}

func (s *APIServer) Serve() {
	fmt.Println("Gateway Server running at: ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, nil))
}

func (s *APIServer) handleGetProducts(w http.ResponseWriter, r *http.Request) error {
	// Parse ids from query params from string[] to int[]
	idsQueryParams := r.URL.Query()["id"]
	var ids []int = nil
	if len(idsQueryParams) != 0 {
		ids = make([]int, len(idsQueryParams))
		for i, idStr := range idsQueryParams {
			var err error
			ids[i], err = strconv.Atoi(idStr)
			if err != nil {
				return internal.NewAPIError(http.StatusBadRequest, "invalid value for query parameter")
			}
		}
	}

	resp, err := s.svc.GetProducts(context.TODO(), ids)
	if err != nil {

	}

	return internal.WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleGetProductById(w http.ResponseWriter, r *http.Request) error {
	id, err := internal.GetPathId("id", r)
	if err != nil {
		return internal.InvalidPathIdentifier()
	}

	resp, err := s.svc.GetProductById(context.TODO(), id)
	if err != nil {
		
	}

	return internal.WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleCreateProduct(w http.ResponseWriter, r *http.Request) error {
	var req internal.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return internal.InvalidRequestBody()
	}

	resp, err := s.svc.CreateProduct(context.TODO(), req)
	if err != nil {
		
	}

	return internal.WriteJSON(w, http.StatusCreated, resp)
}


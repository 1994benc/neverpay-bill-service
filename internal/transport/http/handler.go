package http

import (
	"1994benc/neverpay-api/internal/bill"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router      *mux.Router
	BillService *bill.Service
}

// Creates a new instance of Handler
func New(service *bill.Service) *Handler {
	return &Handler{
		BillService: service,
	}
}

// Setup all routes
func (handler *Handler) SetupRoutes() {
	log.Println("Setting up routes")
	handler.Router = mux.NewRouter()

	// All routes
	handler.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am still alive :D")
	})
	handler.Router.HandleFunc("/api/bills/{id}", handler.GetBill).Methods(http.MethodGet)
	handler.Router.HandleFunc("/api/bills", handler.GetAllBills).Methods(http.MethodGet)
	handler.Router.HandleFunc("/api/bills", handler.AddBill).Methods(http.MethodPost)
	handler.Router.HandleFunc("/api/bills/{id}", handler.DeleteBill).Methods(http.MethodDelete)
	handler.Router.HandleFunc("/api/bills/{id}", handler.UpdateBill).Methods(http.MethodPut)

}

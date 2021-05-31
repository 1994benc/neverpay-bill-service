package http

import (
	"1994benc/neverpay-api/internal/bill"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router      *mux.Router
	BillService *bill.Service
}

// stores responses from API
type GenericResponse struct {
	Message string
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GenericResponse{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
	handler.Router.HandleFunc("/api/bills/{id}", handler.GetBill).Methods(http.MethodGet)
	handler.Router.HandleFunc("/api/bills", handler.GetAllBills).Methods(http.MethodGet)
	handler.Router.HandleFunc("/api/bills", handler.AddBill).Methods(http.MethodPost)
	handler.Router.HandleFunc("/api/bills/{id}", handler.DeleteBill).Methods(http.MethodDelete)
	handler.Router.HandleFunc("/api/bills/{id}", handler.UpdateBill).Methods(http.MethodPut)

}

// Private Methods
func (h *Handler) parseID(idStr string) (uint, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	return uint(id), err
}

func (h *Handler) commonHeadersSetUp(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

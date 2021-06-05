package http

import (
	"1994benc/neverpay-api/internal/bill"
	"1994benc/neverpay-api/internal/user"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router      *mux.Router
	BillService *bill.Service
	UserService *user.Service
}

// stores responses from API
type GenericResponse struct {
	Message string
}

// Creates a new instance of Handler
func New(billService *bill.Service, userService *user.Service) *Handler {
	return &Handler{
		BillService: billService,
		UserService: userService,
	}
}

// Logs out which endpoint has been hit
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method":          r.Method,
				"Endpoint":        r.URL.Path,
				"Route Variables": mux.Vars(r),
			},
		).Info("endpoint hit")
		next.ServeHTTP(w, r)
	})
}

// Protects a route with authentication
func (h *Handler) AuthMiddleware(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Auth endpoint is requested!")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := authHeaderParts[1]

		if h.UserService.ValidateToken(token) {
			original(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	}
}

// Setup all routes
func (handler *Handler) SetupRoutes() {
	log.Println("Setting up routes")
	handler.Router = mux.NewRouter()

	// Middlewares
	handler.Router.Use(LoggingMiddleware)

	// All routes
	handler.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GenericResponse{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
	// User routes
	handler.Router.HandleFunc("/api/user/signup", handler.SignUp).Methods(http.MethodPost)
	handler.Router.HandleFunc("/api/user/signin", handler.SignIn).Methods(http.MethodPost)

	// Bill routes
	handler.Router.HandleFunc("/api/bills/{id}", handler.GetBill).Methods(http.MethodGet)
	handler.Router.HandleFunc("/api/bills", handler.GetAllBills).Methods(http.MethodGet)
	handler.Router.HandleFunc("/api/bills", handler.AddBill).Methods(http.MethodPost)
	handler.Router.HandleFunc("/api/bills/{id}", handler.AuthMiddleware(handler.DeleteBill)).Methods(http.MethodDelete)
	handler.Router.HandleFunc("/api/bills/{id}", handler.AuthMiddleware(handler.UpdateBill)).Methods(http.MethodPut)

}

// Private Methods
func (h *Handler) parseID(idStr string) (uint, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	return uint(id), err
}

func (h *Handler) commonHeadersSetUp(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

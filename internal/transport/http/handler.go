package http

import (
	"1994benc/neverpay-api/internal/bill"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	jwt "github.com/dgrijalva/jwt-go"
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

func validateToken(accessToken string) bool {
	log.Printf("Validating access token %s", accessToken)
	// replace this by loading in a private RSA cert for more security
	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error parsing access token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		log.Error(err)
		return false
	}

	return token.Valid
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

func Auth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
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
		if validateToken(token) {
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
	handler.Router.Use(LoggingMiddleware)

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
	handler.Router.HandleFunc("/api/bills", Auth(handler.AddBill)).Methods(http.MethodPost)
	handler.Router.HandleFunc("/api/bills/{id}", Auth(handler.DeleteBill)).Methods(http.MethodDelete)
	handler.Router.HandleFunc("/api/bills/{id}", Auth(handler.UpdateBill)).Methods(http.MethodPut)

}

// Private Methods
func (h *Handler) parseID(idStr string) (uint, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	return uint(id), err
}

func (h *Handler) commonHeadersSetUp(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

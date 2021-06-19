package http

import (
	"1994benc/neverpay-user-service/internal/user"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router      *mux.Router
	UserService *user.Service
}

// Creates a new instance of Handler
func New(userService *user.Service) *Handler {
	return &Handler{
		UserService: userService,
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
	handler.Router.HandleFunc("/api/users", handler.GetUsers).Methods(http.MethodGet)
	handler.Router.HandleFunc("/api/users/signup", handler.SignUp).Methods(http.MethodPost)
	handler.Router.HandleFunc("/api/users/signin", handler.SignIn).Methods(http.MethodPost)
}

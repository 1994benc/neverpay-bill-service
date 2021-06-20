package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/1994benc/neverpay-bill-service/internal/transport/http/middleware"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

// Creates a new instance of Handler
// TODO: pass in an instance of a service
func New() *Handler {
	// TODO: return &Handler{
	// 	UserService: userService,
	// }
	return &Handler{}
}

// Setup all routes
func (handler *Handler) SetupRoutes() {
	log.Println("Setting up routes")
	handler.Router = mux.NewRouter()

	// Middlewares
	handler.Router.Use(middleware.LoggingMiddleware)

	// All routes
	handler.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GenericResponse{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
}

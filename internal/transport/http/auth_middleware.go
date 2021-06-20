package http

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

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

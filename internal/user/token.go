package user

import (
	"encoding/json"
	"net/http"
)

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func (t *Token) ToJSON(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(t)
}

package user

import (
	"encoding/json"
	"io"
	"net/http"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Authentication) FromJSON(body io.ReadCloser) error {
	err := json.NewDecoder(body).Decode(a)
	return err
}

func (a *Authentication) ToJSON(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(a)
}

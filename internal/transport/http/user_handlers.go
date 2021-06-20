package http

import (
	"1994benc/neverpay-api/internal/user"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := u.FromJSON(r.Body)
	if err != nil {
		log.Error("Error parsing the JSON body: %s", err)
		http.Error(w, "Error parsing the JSON body", http.StatusBadRequest)
		return
	}
	userAlreadyExists := h.checkIfUserExists(u)
	if userAlreadyExists {
		http.Error(w, "Email already in use!", http.StatusBadRequest)
		return
	}

	err = u.HashPassword()
	if err != nil {
		http.Error(w, "Error generating hashed password", http.StatusInternalServerError)
		return
	}

	newUser, err := h.UserService.CreateUser(u)
	if err != nil {
		http.Error(w, "Error creating user! "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = newUser.ToJSON(w)
	if err != nil {
		http.Error(w, "Error writing to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var authDetails user.Authentication
	error := authDetails.FromJSON(r.Body)
	if error != nil {
		http.Error(w, "Error parsing inputs", http.StatusBadRequest)
		return
	}

	u, err := h.UserService.FindUserByEmail(authDetails.Email)
	if u.Email == "" || err != nil {
		http.Error(w, "User not found in our system!", http.StatusForbidden)
		return
	}

	passwordsMatched := u.CheckPasswordHash(authDetails.Password)
	if !passwordsMatched {
		http.Error(w, "Password entered is incorrect!", http.StatusForbidden)
		return
	}

	validToken, err := h.UserService.GenerateJWT(u.Email, "basic")
	if err != nil {
		http.Error(w, "Error generating access token!", http.StatusInternalServerError)
		return
	}

	var token user.Token
	token.Email = u.Email
	token.Role = u.Role
	token.TokenString = validToken
	err = token.ToJSON(w)
	if err != nil {
		http.Error(w, "Error parsing generated token!", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) checkIfUserExists(userInstance user.User) bool {
	u, err := h.UserService.FindUserByEmail(userInstance.Email)
	log.Printf("checkIfUserExists: %s", u.Email)
	return err == nil && u.Email != ""
}

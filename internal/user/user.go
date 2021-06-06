package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserNoPassword struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
	Role  string `json:"role"`
}

func (userNoPassword *UserNoPassword) fromUser(u *User) {
	userNoPassword.ID = u.ID
	userNoPassword.CreatedAt = u.CreatedAt
	userNoPassword.DeletedAt = u.DeletedAt
	userNoPassword.Email = u.Email
	userNoPassword.Model = u.Model
	userNoPassword.Name = u.Name
	userNoPassword.Role = u.Role
	userNoPassword.UpdatedAt = u.UpdatedAt
}

func (u *User) FromJSON(body io.ReadCloser) error {
	err := json.NewDecoder(body).Decode(u)
	return err
}

func (u *User) ToJSON(w http.ResponseWriter) error {
	// Omits the password field
	var userNoPassword UserNoPassword
	userNoPassword.fromUser(u)
	return json.NewEncoder(w).Encode(userNoPassword)
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
	return err
}

func (u *User) CheckPasswordHash(passwordToCheck string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwordToCheck))
	return err == nil
}

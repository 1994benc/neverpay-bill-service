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

func (u *User) FromJSON(body io.ReadCloser) error {
	err := json.NewDecoder(body).Decode(u)
	return err
}

func (u *User) ToJSON(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(u)
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

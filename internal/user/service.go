package user

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// TODO: remove this secret key
const secretkey = "testsecretkeytoberemoved"

type IService interface {
	CreateUser(user User) (User, error)
	FindUserByEmail(email string) (User, error)
}

type Service struct {
	DB *gorm.DB
}

// Returns a new User service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) CreateUser(u User) (User, error) {
	result := s.DB.Save(u)
	return u, result.Error
}

func (s *Service) FindUserByEmail(email string) (User, error) {
	var user User
	result := s.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (s *Service) GenerateJWT(email string, role string) (string, error) {
	var mySigningKey = []byte(secretkey) // TODO: use a secure secretkey
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Printf("Error generating JWT: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func (s *Service) ValidateToken(accessToken string) bool {
	log.Printf("Validating access token %s", accessToken)
	// replace this by loading in a private RSA cert for more security
	var mySigningKey = []byte(secretkey)
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error parsing access token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		log.Println(err)
		return false
	}

	return token.Valid
}

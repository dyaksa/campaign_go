package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(ID int, email string, name string, occupation string) (string, error)
}

type jwtService struct {
}

func NewAuthService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = "s3Cr3T_K3Y_T0k3n"

func (t *jwtService) GenerateToken(ID int, email string, name string, occupation string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = ID
	claims["user_email"] = email
	claims["user_name"] = name
	claims["user_occupation"] = occupation
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return token, err
	}
	return token, nil
}

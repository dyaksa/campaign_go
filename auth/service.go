package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(ID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewAuthService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = "s3Cr3T_K3Y_T0k3n"

func (t *jwtService) GenerateToken(ID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = ID
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return token, err
	}
	return token, nil
}

func (t *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	tokenValid, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token not valid")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return tokenValid, nil
}

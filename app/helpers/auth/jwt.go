package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
	"os"
	"time"
)

type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(720)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	return token.SignedString([]byte(os.Getenv("APP_KEY")))
}

func ValidateRequest(req *http.Request) (CustomClaims, error) {
	var claim CustomClaims
	_, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("APP_KEY"))
		return b, nil
	}, request.WithClaims(&claim))
	return claim, err
}

func GenerateTokenFromClaims(claims CustomClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString([]byte(os.Getenv("APP_KEY")))
}

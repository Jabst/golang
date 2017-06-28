package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SetToken(username string) http.Cookie {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	expireCookie := time.Now().Add(time.Hour * 1)

	claims := Claims {
		username,
		jwt.StandardClaims {
			ExpiresAt	: expireToken,
			Issuer 		: "localhost:8081",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))

	return http.Cookie {
		Name		: "Auth",
		Value		: signedToken,
		Expires		: expireCookie,
		HttpOnly	: true,
	}
}
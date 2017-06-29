package utils

import (
	"net/http"
	"time"
	"log"

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

func ValidateToken(t string) bool {

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error){
		return []byte("secret"), nil
	})

	log.Println(token)

	log.Println(err)

	if err != nil {
		log.Println("Not Found")
		return false
	}

	return true
}
package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("tinyurlservicesecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := JWTClaim{
		Username: "bar",
		Email:    "test",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "tinyurlgolangimplementation",
		},
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"email":    email,
	// 	"username": username,
	// 	"exp":      expirationTime.Unix(),
	// })
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if token.Valid {
		fmt.Printf("Token is valid. Username: %s Expires at: %s", claims.Username, claims.ExpiresAt)
		return nil
	}

	if errors.Is(err, jwt.ErrTokenMalformed) {
		err = errors.New("invalid token format")
		return
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		err = errors.New("token is either expired or not active yet")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	return

}

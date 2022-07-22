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

type ContextInfo struct {
	Username string
	Email    string
	UserID   string
}

func GenerateJWT(userid string, email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := JWTClaim{
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "tinyurlgolangimplementation",
			Subject:   userid,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (*ContextInfo, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	if token.Valid {
		fmt.Printf("Token is valid. Username: %s Expires at: %s", claims.Username, claims.ExpiresAt)
		return &ContextInfo{UserID: claims.Subject, Username: claims.Username, Email: claims.Email}, err
	}

	if errors.Is(err, jwt.ErrTokenMalformed) {
		err = errors.New("invalid token format")
		return nil, err
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		err = errors.New("token is either expired or not active yet")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	return nil, err
}

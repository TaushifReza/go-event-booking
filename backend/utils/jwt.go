package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessSecret  = os.Getenv("JWT_ACCESS_SECRET")
	refreshSecret = os.Getenv("JWT_REFRESH_SECRET")
)

type JWTClaims struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	TokenType string `json:"token_type"` // access / refresh
	jwt.RegisteredClaims
}

func GenerateAccessToken(id uint, email string) (string, error) {
	claims := &JWTClaims{
		ID:        id,
		Email:     email,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-event-api",
			Audience:  []string{"go-event-client"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

func GenerateRefreshToken(id uint, email string) (string, error) {
	claims := &JWTClaims{
		ID:        id,
		Email:     email,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-event-api",
			Audience:  []string{"go-event-client"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)), // 30 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshSecret))
}

func VerifyToken(tokenStr, tokenType string) (*JWTClaims, error) {
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	var secret []byte
	if tokenType == "access" {
		secret = []byte(accessSecret)
	} else {
		secret = []byte(refreshSecret)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		fmt.Println("TOKEN PARSE ERROR:", err)
		return nil, errors.New("could not parse token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Optional: enforce type
	if claims.TokenType != tokenType {
		return nil, errors.New("token type mismatch")
	}

	return claims, nil
}

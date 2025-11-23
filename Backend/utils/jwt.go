package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

func GenerateToken(id uint, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"email": email,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})	

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (uint, error){
	token = strings.TrimPrefix(token, "Bearer ")
	
	parseToken, err := jwt.Parse(token, func (token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok{
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("PARSE ERROR: ",err)
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parseToken.Valid
	
	if !tokenIsValid {
		return 0,errors.New("invalid token")
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	id := uint(claims["id"].(float64))

	return id, nil
}
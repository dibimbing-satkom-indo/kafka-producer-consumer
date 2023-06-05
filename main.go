package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	jwt.RegisteredClaims
	ID string `json:"id_user"`
}

func GenerateJWT(id, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateUserJWT(id, secret string) (string, error) {
	claims := User{ID: id}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJWT(jwtToken, secret string) (string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data, _ := json.Marshal(claims)
		user := User{}
		_ = json.Unmarshal(data, &user)

		return fmt.Sprintf("%s", user.ID), nil
	}

	return "", errors.New("invalid token")
}

func main() {
	secret := "rahasia"
	jwtToken, _ := GenerateUserJWT("10", secret)
	fmt.Println("jwtToken", jwtToken)
	id, _ := VerifyJWT(jwtToken, secret)
	fmt.Println("id", id)
}

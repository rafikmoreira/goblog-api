package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/rafikmoreira/go-blog-api/cmd/API/config"
)

func VerifyJWT(token string) bool {
	if token == "" {
		return false
	}

	tknStr := token

	tkn, err := jwt.ParseWithClaims(tknStr, nil, func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	})

	fmt.Print(tkn)

	if err != nil {
		return false
	}

	if !tkn.Valid {
		return false
	}
	return true
}

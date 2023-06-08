package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/rafikmoreira/go-blog-api/cmd/API/config"
)

func VerifyJWT(token string) bool {
	if len(token) == 0 {
		return false
	}

	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unauthorized")
		}
		return config.SecretKey, nil
	})

	if err != nil {
		return false
	}

	if !tkn.Valid {
		return false
	}

	return true
}

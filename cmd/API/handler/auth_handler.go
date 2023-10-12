package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rafikmoreira/go-blog-api/cmd/API/config"
	"github.com/rafikmoreira/go-blog-api/internal/entity"
	"net/http"
	"time"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	UserUseCase *entity.IUserUseCase
}

func (h AuthHandler) Login(c *gin.Context) {
	user := new(entity.User)

	userUseCase := *h.UserUseCase
	credentials := new(Credentials)

	err := c.BindJSON(credentials)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "e-mail or password incorrect",
		})
		return
	}

	user, err = userUseCase.GetByEmail(&credentials.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "e-mail or password incorrect",
		})
		return
	}

	if !entity.CheckPasswordHash(&credentials.Password, &user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "e-mail or password incorrect",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := jwt.MapClaims{
		"expires_at": time.Now().Add(time.Hour).Unix(),
		"authorized": true,
		"user":       user,
	}
	token.Claims = claims

	tokenString, err := token.SignedString(config.SecretKey)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

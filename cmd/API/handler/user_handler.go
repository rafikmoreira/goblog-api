package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/domain"
	"net/http"
)

type UserHandler struct {
	UserUseCase *domain.IUserUseCase
}

func (h UserHandler) CreateUser(c *gin.Context) {
	user := new(domain.User)
	userUseCase := *h.UserUseCase

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = userUseCase.Create(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível cadastrar o usuário",
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h UserHandler) ListUser(c *gin.Context) {
	userUseCase := *h.UserUseCase
	users, err := userUseCase.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar os usuários",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h UserHandler) GetUser(c *gin.Context) {
	id := c.Param("userId")
	userUseCase := *h.UserUseCase
	user, err := userUseCase.GetByID(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h UserHandler) UpdateUser(c *gin.Context) {
	userUseCase := *h.UserUseCase
	data := &domain.User{}
	err := c.ShouldBindJSON(data)
	id := c.Param("userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userUseCase.Update(data, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível atualizar os dados do usuário",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h UserHandler) DeleteUser(c *gin.Context) {
	userUseCase := *h.UserUseCase
	id := c.Param("userId")
	err := userUseCase.Destroy(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar os usuários",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "usuário deletado com sucesso",
	})
}

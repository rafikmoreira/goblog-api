package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/application"
	"github.com/rafikmoreira/go-blog-api/domain"
	"net/http"
)

func CreateUserHandler(c *gin.Context, repository *domain.IUserRepository) {
	user := new(domain.User)

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = application.UserUseCaseImplementation.Create(repository, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível cadastrar o usuário",
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func ListUserHandler(c *gin.Context, repository *domain.IUserRepository) {
	users, err := application.UserUseCaseImplementation.List(repository)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar os usuários",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserHandler(c *gin.Context, repository *domain.IUserRepository) {
	id := c.Param("userId")
	user, err := application.UserUseCaseImplementation.GetByID(repository, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context, repository *domain.IUserRepository) {
	data := &domain.User{}
	err := c.ShouldBindJSON(data)
	id := c.Param("userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := application.UserUseCaseImplementation.Update(repository, data, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível atualizar os dados do usuário",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUserHandler(c *gin.Context, repository *domain.IUserRepository) {
	id := c.Param("userId")
	err := application.UserUseCaseImplementation.Destroy(repository, &id)

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

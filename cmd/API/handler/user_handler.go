package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/application"
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/repository"
	"net/http"
)

func CreateUserHandler(c *gin.Context) {
	user := new(domain.User)

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = application.UserUseCaseImplementation.Create(repository.UserRepositoryImplementation, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível cadastrar o usuário",
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func ListUserHandler(c *gin.Context) {
	users, err := application.UserUseCaseImplementation.List(repository.UserRepositoryImplementation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar os usuários",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserHandler(c *gin.Context) {
	id := c.Param("userId")
	user, err := application.UserUseCaseImplementation.GetByID(repository.UserRepositoryImplementation, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context) {
	data := &domain.User{}
	err := c.ShouldBindJSON(data)
	id := c.Param("userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := application.UserUseCaseImplementation.Update(repository.UserRepositoryImplementation, data, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar os usuários",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("userId")
	err := application.UserUseCaseImplementation.Destroy(repository.UserRepositoryImplementation, &id)

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

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/application"
	"github.com/rafikmoreira/go-blog-api/domain"
)

func CreatePostHandler(c *gin.Context, repository *domain.IPostRepository) {
	post := new(domain.Post)

	err := c.ShouldBindJSON(&post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = application.PostUseCaseImplementation.Create(repository, post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível cadastrar o post",
		})
		return
	}
	c.JSON(http.StatusCreated, post)
}
func ListPostHandler(c *gin.Context, repository *domain.IPostRepository) {
	posts, err := application.PostUseCaseImplementation.List(repository)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar as postagens",
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}
func GetPostHandler(c *gin.Context, repository *domain.IPostRepository) {
	id := c.Param("postId")
	post, err := application.PostUseCaseImplementation.GetByID(repository, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "post não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
func UpdatePostHandler(c *gin.Context, repository *domain.IPostRepository) {
	data := &domain.Post{}
	err := c.ShouldBindJSON(data)
	id := c.Param("postId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := application.PostUseCaseImplementation.Update(repository, data, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar as postagens",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
func DeletePostHandler(c *gin.Context, repository *domain.IPostRepository) {
	id := c.Param("postId")
	err := application.PostUseCaseImplementation.Destroy(repository, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar as postagens",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "post deletado com sucesso",
	})
}

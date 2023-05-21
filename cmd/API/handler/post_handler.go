package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/application"
	"github.com/rafikmoreira/go-blog-api/domain"
	"net/http"
)

func CreatePostHandler(c *gin.Context) {
	post := new(domain.Post)

	err := c.ShouldBindJSON(&post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = application.NewPostUseCaseImplementation.Create(post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível cadastrar o post",
		})
		return
	}
	c.JSON(http.StatusCreated, post)
}
func ListPostHandler(c *gin.Context) {
	posts, err := application.NewPostUseCaseImplementation.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar as postagens",
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}
func GetPostHandler(c *gin.Context) {
	id := c.Param("postId")
	post, err := application.NewPostUseCaseImplementation.GetByID(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "post não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
func UpdatePostHandler(c *gin.Context) {
	data := &domain.Post{}
	err := c.ShouldBindJSON(data)
	id := c.Param("postId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := application.NewPostUseCaseImplementation.Update(data, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar as postagens",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
func DeletePostHandler(c *gin.Context) {
	id := c.Param("postId")
	err := application.NewPostUseCaseImplementation.Destroy(&id)

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

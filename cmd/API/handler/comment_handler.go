package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/application"
	"github.com/rafikmoreira/go-blog-api/domain"
	"net/http"
)

func CreateCommentHandler(c *gin.Context) {
	comment := new(domain.Comment)
	postId := c.Param("postId")
	err := c.ShouldBindJSON(&comment)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = application.NewCommentUseCaseImplementation.Create(comment, &postId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível adicionar o comentário",
		})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

func DeleteCommentHandler(c *gin.Context) {
	commentId := c.Param("commentId")
	postId := c.Param("postId")

	err := application.NewCommentUseCaseImplementation.Destroy(&postId, &commentId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar os comentários",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "comment deletado com sucesso",
	})
}

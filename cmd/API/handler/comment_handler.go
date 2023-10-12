package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/internal/entity"
)

type CommentHandler struct {
	CommentUseCase *entity.ICommentUseCase
}

func (h CommentHandler) CreateComment(c *gin.Context) {
	commentUseCase := *h.CommentUseCase
	comment := new(entity.Comment)
	postId := c.Param("postId")
	err := c.ShouldBindJSON(&comment)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = commentUseCase.Create(comment, &postId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível adicionar o comentário",
		})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

func (h CommentHandler) DeleteComment(c *gin.Context) {
	commentUseCase := *h.CommentUseCase
	commentId := c.Param("commentId")
	postId := c.Param("postId")

	err := commentUseCase.Destroy(&postId, &commentId)

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

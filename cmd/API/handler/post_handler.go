package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/domain"
)

type PostHandler struct {
	PostUseCase *domain.IPostUseCase
}

func (h PostHandler) CreatePost(c *gin.Context) {
	post := new(domain.Post)
	postUseCase := *h.PostUseCase
	err := c.ShouldBindJSON(&post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postUseCase.Create(post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível cadastrar o post",
		})
		return
	}
	c.JSON(http.StatusCreated, post)
}
func (h PostHandler) ListPost(c *gin.Context) {
	postUseCase := *h.PostUseCase
	posts, err := postUseCase.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar as postagens",
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}
func (h PostHandler) GetPost(c *gin.Context) {
	id := c.Param("postId")
	postUseCase := *h.PostUseCase
	post, err := postUseCase.GetByID(&id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "post não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
func (h PostHandler) UpdatePost(c *gin.Context) {
	postUseCase := *h.PostUseCase
	data := &domain.Post{}
	err := c.ShouldBindJSON(data)
	id := c.Param("postId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := postUseCase.Update(data, &id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não foi possível listar as postagens",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
func (h PostHandler) DeletePost(c *gin.Context) {
	postUseCase := *h.PostUseCase
	id := c.Param("postId")
	err := postUseCase.Destroy(&id)

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

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/cmd/API/handler"
	"github.com/rafikmoreira/go-blog-api/cmd/API/middleware"
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/repository"
)

func main() {
	// repositories
	commentRepository := domain.NewCommentRepository(new(repository.CommentRepository))
	postRepository := domain.NewPostRepository(new(repository.PostRepository))
	userRepository := domain.NewUserRepository(new(repository.UserRepository))

	// gin router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Welcome to the Go Blog API",
			"author":   "Rafik Moreira",
			"github":   "https://github.com/rafikmoreira",
			"linkedin": "https://www.linkedin.com/in/rafikmoreira/",
		})
	})

	// post routes
	r.POST("/post", middleware.AuthMiddleware(), func(context *gin.Context) {
		handler.CreatePostHandler(context, postRepository)
	})
	r.GET("/post", func(context *gin.Context) {
		handler.ListPostHandler(context, postRepository)
	})
	r.GET("/post/:postId", func(context *gin.Context) {
		handler.GetPostHandler(context, postRepository)
	})
	r.PATCH("/post/:postId", middleware.AuthMiddleware(), func(context *gin.Context) {
		handler.UpdatePostHandler(context, postRepository)
	})
	r.DELETE("/post/:postId", middleware.AuthMiddleware(), func(context *gin.Context) {
		handler.DeletePostHandler(context, postRepository)
	})

	// comment routes
	r.POST("/post/:postId/comment", middleware.AuthMiddleware(), func(context *gin.Context) {
		handler.CreateCommentHandler(context, commentRepository)
	})
	r.DELETE("/post/:postId/comment/:commentId", middleware.AuthMiddleware(), func(context *gin.Context) {
		handler.DeleteCommentHandler(context, commentRepository)
	})

	// user routes
	r.POST("/user", func(context *gin.Context) {
		handler.CreateUserHandler(context, userRepository)
	})
	r.GET("/user", func(context *gin.Context) {
		handler.ListUserHandler(context, userRepository)
	})
	r.GET("/user/:userId", func(context *gin.Context) {
		handler.GetUserHandler(context, userRepository)
	})
	r.PATCH("/user/:userId", middleware.AuthMiddleware(), func(context *gin.Context) {
		handler.UpdateUserHandler(context, userRepository)
	})
	r.DELETE("/user/:userId", middleware.AuthMiddleware(), func(context *gin.Context) {
		handler.DeleteUserHandler(context, userRepository)
	})

	// auth routes
	r.POST("/auth", func(context *gin.Context) {
		handler.AuthHandler(context, userRepository)
	})

	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	err = r.Run(":3333")

	if err != nil {
		panic(err)
	}
}

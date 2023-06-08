package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/cmd/API/handler"
	"github.com/rafikmoreira/go-blog-api/cmd/API/middleware"
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/db"
	"github.com/rafikmoreira/go-blog-api/infrastructure/repository"
	"github.com/rafikmoreira/go-blog-api/use_case"
)

func main() {
	// db connection
	dbConnection := db.NewSQLiteDBConnection(db.Migrate)

	// repositories
	var commentRepository domain.ICommentRepository = &repository.CommentRepository{
		DB: dbConnection,
	}
	var postRepository domain.IPostRepository = &repository.PostRepository{
		DB: dbConnection,
	}
	var userRepository domain.IUserRepository = &repository.UserRepository{
		DB: dbConnection,
	}

	// use cases
	var commentUseCase domain.ICommentUseCase = &use_case.CommentUseCase{
		Repository: &commentRepository,
	}
	var postUseCase domain.IPostUseCase = &use_case.PostUseCase{
		Repository: &postRepository,
	}
	var userUseCase domain.IUserUseCase = &use_case.UserUseCase{
		Repository: &userRepository,
	}

	// handlers
	commentHandler := &handler.CommentHandler{
		CommentUseCase: &commentUseCase,
	}
	postHandler := &handler.PostHandler{
		PostUseCase: &postUseCase,
	}
	authHandler := &handler.AuthHandler{
		UserUseCase: &userUseCase,
	}
	userHandler := &handler.UserHandler{
		UserUseCase: &userUseCase,
	}

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
	r.POST("/post", middleware.AuthMiddleware(), postHandler.CreatePost)
	r.GET("/post", postHandler.ListPost)
	r.GET("/post/:postId", postHandler.GetPost)
	r.PATCH("/post/:postId", middleware.AuthMiddleware(), postHandler.UpdatePost)
	r.DELETE("/post/:postId", middleware.AuthMiddleware(), postHandler.DeletePost)

	// comment routes
	r.POST("/post/:postId/comment", middleware.AuthMiddleware(), commentHandler.CreateComment)
	r.DELETE("/post/:postId/comment/:commentId", middleware.AuthMiddleware(), commentHandler.DeleteComment)

	// user routes
	r.POST("/user", userHandler.CreateUser)
	r.GET("/user", userHandler.ListUser)
	r.GET("/user/:userId", userHandler.GetUser)
	r.PATCH("/user/:userId", middleware.AuthMiddleware(), userHandler.UpdateUser)
	r.DELETE("/user/:userId", middleware.AuthMiddleware(), userHandler.DeleteUser)

	// auth routes
	r.POST("/auth", authHandler.Login)

	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	err = r.Run(":3333")

	if err != nil {
		panic(err)
	}
}

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
	dbConnection := db.NewPostgreSQLConnection(db.Migrate)

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
	postRoutes := r.Group("/post")
	postRoutes.POST("", middleware.AuthMiddleware(), postHandler.CreatePost)
	postRoutes.GET("", postHandler.ListPost)
	postRoutes.GET("/:postId", postHandler.GetPost)
	postRoutes.PATCH("/:postId", middleware.AuthMiddleware(), postHandler.UpdatePost)
	postRoutes.DELETE("/:postId", middleware.AuthMiddleware(), postHandler.DeletePost)
	postRoutes.POST("/:postId/comment", middleware.AuthMiddleware(), commentHandler.CreateComment)
	postRoutes.DELETE("/:postId/comment/:commentId", middleware.AuthMiddleware(), commentHandler.DeleteComment)

	// user routes
	userRoutes := r.Group("/user")
	userRoutes.POST("", userHandler.CreateUser)
	userRoutes.GET("", userHandler.ListUser)
	userRoutes.GET("/:userId", userHandler.GetUser)
	userRoutes.PATCH("/:userId", middleware.AuthMiddleware(), userHandler.UpdateUser)
	userRoutes.DELETE("/:userId", middleware.AuthMiddleware(), userHandler.DeleteUser)

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

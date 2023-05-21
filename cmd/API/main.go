package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/cmd/API/handler"
)

func main() {
	r := gin.Default()

	// Post routes
	r.POST("/post", handler.CreatePostHandler)
	r.GET("/post", handler.ListPostHandler)
	r.GET("/post/:id", handler.GetPostHandler)
	r.PATCH("/post/:id", handler.UpdatePostHandler)
	r.DELETE("/post/:id", handler.DeletePostHandler)

	// User routes
	r.POST("/user", handler.CreateUserHandler)
	r.GET("/user", handler.ListUserHandler)
	r.GET("/user/:id", handler.GetUserHandler)
	r.PATCH("/user/:id", handler.UpdateUserHandler)
	r.DELETE("/user/:id", handler.DeleteUserHandler)

	err := r.Run(":3333")

	if err != nil {
		panic(err)
	}
}

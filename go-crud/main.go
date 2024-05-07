package main

import (
	"github.com/MaaHiN15/go-practice/go-crud/controllers"
	"github.com/MaaHiN15/go-practice/go-crud/initializers"
	"github.com/gin-gonic/gin"
)


func init() {
	initializers.LoadEnvVar()
	initializers.DatabaseSetUp()
	initializers.SyncDB()
}

func main() {
	route := gin.Default()

	// Checking connection
	route.GET("/", controllers.PingPong)

	// Post CRUD operations
	route.GET("/post/:id", controllers.PostShow)
	route.POST("/post", controllers.PostCreate)
	route.PUT("/post/:id", controllers.PostUpdate)
	route.DELETE("/post/:id", controllers.PostDelete)

	// Get all posts
	route.GET("/posts", controllers.GetPosts)

	route.Run()
}
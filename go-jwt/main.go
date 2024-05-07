package main

import (
	"github.com/MaaHiN15/go-practice/go-jwt/controllers"
	"github.com/MaaHiN15/go-practice/go-jwt/initializers"
	"github.com/MaaHiN15/go-practice/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.DatabaseSetUp()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.PingPong)
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	
	r.Run()
}
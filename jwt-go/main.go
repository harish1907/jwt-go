package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harish1907/jwt-go/controllers"
	"github.com/harish1907/jwt-go/intializers"
	"github.com/harish1907/jwt-go/middleware"
)

func init() {
	intializers.LocalEnvironmentVariable()
	intializers.ConnectionDB()
	intializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequriedAuth, controllers.Validate)
	r.Run()
}

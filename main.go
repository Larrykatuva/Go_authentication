package main

import (
	"github.com/gin-gonic/gin"
	"go-authentication/controllers"
	"go-authentication/initializers"
	"go-authentication/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.User{})
}

func main() {
	r := gin.Default()
	r.POST("register", controllers.SignUpUser)
	r.POST("login", controllers.SignInUser)
	r.GET("users", controllers.GetUsers)
	r.Run() // listen and serve on 0.0.0.0:8080
}

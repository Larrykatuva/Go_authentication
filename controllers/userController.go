package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-authentication/dtos"
	"go-authentication/models"
	"go-authentication/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUpUser(c *gin.Context) {
	var body dtos.SignUpUserDto
	c.Bind(&body)
	if user, _ := services.FilterUser(models.User{Email: body.Email}); user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User matching email already exists"})
		return
	}
	if _, err := services.CreateUser(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Verification code send to your email"})
}

func SignInUser(c *gin.Context) {
	var body dtos.SingInUserDto
	c.Bind(&body)
	var user *models.User
	if user, _ = services.FilterUser(models.User{Email: body.Email}); user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User matching the email does not exists"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid login details"})
		return
	}
	fmt.Println(user)
}

func GetUsers(c *gin.Context) {
	var users *[]models.User
	users, _ = services.FilterUsers(models.User{})
	c.JSON(http.StatusOK, users)
}

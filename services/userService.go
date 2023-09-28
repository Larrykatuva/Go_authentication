package services

import (
	"go-authentication/dtos"
	"go-authentication/initializers"
	"go-authentication/models"
	"go-authentication/shared"
	"golang.org/x/crypto/bcrypt"
)

func generatePasswordHash(password string) (hashedPassword string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func CreateUser(userDto dtos.SignUpUserDto) (*models.User, error) {
	user := models.User{Email: userDto.Email, Username: userDto.UserName, Password: generatePasswordHash(userDto.Password), Inactive: false, Verified: false}
	if err := initializers.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FilterUser(filter models.User) (*models.User, error) {
	var user models.User
	result := initializers.DB.Where(filter).First(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FilterUsers(filters models.User) (*[]models.User, error) {
	var users []models.User
	result := initializers.DB.Where(filters).Find(&users)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func FilterPaginatedUsers(filters models.User, pagination shared.Pagination) (*[]models.User, error) {
	var users []models.User
	result := initializers.DB.Where(filters).Offset(pagination.Offset).Limit(pagination.Limit).Find(&users)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

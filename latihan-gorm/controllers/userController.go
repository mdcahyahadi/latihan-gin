package controllers

import (
	"errors"
	"fmt"
	"latihan-gorm/database"
	"latihan-gorm/models"

	"gorm.io/gorm"
)

func CreateUser(email string) {
	db := database.GetDB()
	User := models.User{
		Email: email,
	}
	err := db.Create(&User).Error
	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}
	fmt.Println("New user: ", User)
}
func GetAllUserWithProduct() {
	db := database.GetDB()
	users := models.User{}
	err := db.Preload("Products").Find(&users).Error
	if err != nil {
		fmt.Println("Error getting users with product: ", err)
		return
	}
	fmt.Println("Users with product: ")
	fmt.Printf("%v", users)
}
func GetUserById(id uint) {
	db := database.GetDB()
	user := models.User{}
	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found ")
			return
		}
		print("Error finding user:", err)
	}
	fmt.Printf("User Data: %v \n", user)
}
func UpdateUserById(id uint, email string) {
	db := database.GetDB()
	user := models.User{}
	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	// err := db.Updates(models.User{Email: email}).Where("id = ?", id).Error
	if err != nil {
		fmt.Println("Error updatiing user data", err)
		return
	}
	fmt.Printf("Update user's email %v \n", user.Email)
}

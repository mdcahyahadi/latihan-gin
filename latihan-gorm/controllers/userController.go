package controllers

import (
	"fmt"
	"latihan-gorm/database"
	"latihan-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary create user
// @Description create new user data
// @Tags users
// @Accept json
// @Produce json
// @Parameters models.User
// @Success 200 {object} models.User
// @Router /user/new-user [post]
func CreateUser(c *gin.Context) {
	var body struct {
		email string
	}

	db := database.GetDB()
	User := models.User{
		Email: body.email,
	}
	c.Bind(&User)

	err := db.Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": User,
	})
}

// GetAllUserWithProduct godoc
// @Summary get all user with product
// @Description get user with it's coresponding products
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/user-product [get]
func GetAllUserWithProduct(c *gin.Context) {
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

// GetUserById godoc
// @Summary get user by id
// @Description get user with certain id
// @Tags users
// @Accept json
// @Produce json
// @Parameters id
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDB()
	user := models.User{}
	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {

		// }
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// UpdateUserById godoc
// @Summary update user by id
// @Description update user with certain id
// @Tags users
// @Accept json
// @Produce json
// @Parameters id
// @Success 200 {object} models.User
// @Router /user/{id} [patch]
func UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		email string
	}
	db := database.GetDB()
	c.Bind(&body)
	user := models.User{}
	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: body.email}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update user's email" + body.email + "\n",
	})
}

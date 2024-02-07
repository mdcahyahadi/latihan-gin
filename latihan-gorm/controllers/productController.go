package controllers

import (
	"fmt"
	"latihan-gorm/database"
	"latihan-gorm/models"
)

func CreateProduct(userId uint, brand string, name string) {
	db := database.GetDB()
	Product := models.Product{
		UserID: userId,
		BRand:  brand,
		Name:   name,
	}
	err := db.Create(&Product).Error
	if err != nil {
		fmt.Println("Error creating product: ", err)
		return
	}
	fmt.Println("New product: ", Product)
}
func DeleteProduct(id uint) {
	db := database.GetDB()
	product := models.Product{}
	err := db.Delete(&product).Where("id = ?", id).Error
	if err != nil {
		fmt.Println("Error deleting product: ", err)
		return
	}
	fmt.Printf("Product id %v successfully deleted", id)
}

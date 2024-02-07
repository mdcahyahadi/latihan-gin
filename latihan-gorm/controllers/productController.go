package controllers

import (
	"latihan-gorm/database"
	"latihan-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary create product
// @Description create new product data
// @Tags products
// @Accept json
// @Produce json
// @Parameters models.Product
// @Success 200 {object} models.Product
// @Router /product/new-product [post]
func CreateProduct(c *gin.Context) {
	var body struct {
		userId uint
		brand  string
		name   string
	}
	db := database.GetDB()
	Product := models.Product{
		UserID: body.userId,
		BRand:  body.brand,
		Name:   body.name,
	}
	c.Bind(&Product)
	err := db.Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": Product,
	})
}

// DeleteProduct godoc
// @Summary delete product
// @Description delete product by id
// @Tags products
// @Accept json
// @Produce json
// @Parameters models.Product
// @Success 200 {object} models.Product
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	product := models.Product{}
	err := db.Delete(&product).Where("id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data successfully deleted",
	})
}

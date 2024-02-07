package routers

import (
	"latihan-pertama/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/car/new-car", controllers.CreateCar)
	router.PUT("/car/:carID", controllers.UpdateCar)
	router.GET("/car/:carID", controllers.GetCarByID)
	router.GET("/car/", controllers.GetCarAll)
	router.DELETE("/car/:carID", controllers.DeleteCar)
	return router
}

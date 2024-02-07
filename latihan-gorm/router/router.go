package router

import (
	"latihan-gorm/controllers"

	_ "latihan-gorm/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User and Product API
// @version 1.0
// @description This is a service for managing users and products
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	// user
	router.POST("/user/new-user", controllers.CreateUser)
	router.GET("/user/user-product", controllers.GetAllUserWithProduct)
	router.GET("/user/{id}", controllers.GetUserById)
	router.PUT("/user/{id}", controllers.UpdateUserById)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// product
	router.POST("/product/new-product", controllers.CreateProduct)
	router.DELETE("/product/{id}", controllers.DeleteProduct)

	return router
}

package main

import (
	"latihan-gorm/database"
	"latihan-gorm/router"
)

var PORT = ":8080"

func main() {
	database.StartConnection()
	// controllers.CreateUser("mcahyahadi@gmail.com")
	// controllers.CreateProduct(1, "Aqua", "Air Mineral")
	// controllers.GetAllUserWithProduct()
	// controllers.UpdateUserById(2, "rosahey")
	router.StartServer().Run(PORT)
}

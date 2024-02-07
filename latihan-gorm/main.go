package main

import (
	"latihan-gorm/controllers"
	"latihan-gorm/database"
)

func main() {
	database.StartConnection()
	controllers.CreateUser("mcahyahadi@gmail.com")
	// controllers.CreateProduct(1, "Aqua", "Air Mineral")
	// controllers.GetAllUserWithProduct()
	// controllers.UpdateUserById(2, "rosahey")
}

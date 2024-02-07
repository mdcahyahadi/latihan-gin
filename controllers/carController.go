package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json: "car_id"`
	Brand string `json: "brand"`
	Model string `json: "model"`
	Price int    `json: "price"`
}

var cars = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("%d", len(cars)+1)
	cars = append(cars, newCar)
	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}
func UpdateCar(ctx *gin.Context) {
	carId := ctx.Param("carID")
	found := false
	var updatedCar Car
	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for i, car := range cars {
		if carId == car.CarID {
			found = true
			cars[i] = updatedCar
			cars[i].CarID = carId
			break
		}
	}
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carId),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data successfully updated",
		"car":     updatedCar,
	})
}
func GetCarByID(ctx *gin.Context) {
	carId := ctx.Param("carID")
	found := false
	var foundCars []Car
	for i, car := range cars {
		if carId == car.CarID {
			found = true
			foundCars = append(foundCars, cars[i])
			break
		}
	}
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carId),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data successfully fetched",
		"car":     foundCars,
	})
}
func GetCarAll(ctx *gin.Context) {
	var foundCars []Car = cars
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data successfully fetched",
		"car":     foundCars,
	})
}
func DeleteCar(ctx *gin.Context) {
	carId := ctx.Param("carID")
	found := false
	if carId == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "ID Param Empty",
			"error_message": "Please input the car ID you want to delete",
		})
		return
	}
	for i, car := range cars {
		if carId == car.CarID {
			found = true
			copy(cars[i:], cars[i+1:])
			cars[len(cars)-1] = Car{}
			cars = cars[:len(cars)-1]
			break
		}
	}
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carId),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data successfully deleted",
	})
}

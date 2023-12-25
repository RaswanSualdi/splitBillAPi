package main

import (
	"SplitBillApi/configuration"
	"SplitBillApi/controller"
	"SplitBillApi/models"
	"SplitBillApi/repositories"
	"SplitBillApi/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configuration.Connection()
	db.AutoMigrate(&models.Merchant{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.SplitBill{})
	db.AutoMigrate(&models.Customer{})
	repository := repositories.NewRepository(db)
	usecase := usecase.NewUseCase(repository)
	controller := controller.NewController(usecase)
	router := gin.Default()
	api := router.Group("/api")
	api.POST("/splitbill", controller.SplitBill)
	api.GET("/mysplitbill/:id", controller.GetSplitBillByCustomerID)
	router.Run()

}

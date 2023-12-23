package main

import (
	"SplitBillApi/configuration"
	"SplitBillApi/models"
)

func main() {
	db := configuration.Connection()
	db.AutoMigrate(&models.Merchant{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.SplitBill{})
	db.AutoMigrate(&models.Customer{})
}

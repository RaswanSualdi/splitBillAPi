package repositories

import (
	"SplitBillApi/models"

	"gorm.io/gorm"
)

type Repositories interface {
	GetProductById(id int) (models.Product, error)
	GetCustomerById(id int) (models.Customer, error)
	SaveSplitBill(splitBill models.SplitBill) (models.SplitBill, error)
	GetSplitBillByCustomerID(id string) ([]models.SplitBill, error)
}

type repositories struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositories {
	return &repositories{db}
}

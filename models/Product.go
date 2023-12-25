package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ID          int
	SplitBills  []SplitBill `gorm:"many2many:splitbill_products"`
	MerchantID  *string
	ProductName string
	Price       int64
}

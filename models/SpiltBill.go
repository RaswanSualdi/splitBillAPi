package models

import (
	"github.com/jinzhu/gorm"
)

type SplitBill struct {
	gorm.Model
	ID          int
	Customers   []Customer `gorm:"many2many:splitbill_customers;"`
	Products    []Product  `gorm:"many2many:splitbill_products"`
	Description string
	TotalPrice  int64
	TotalAmount int64
	IsEnough    bool
	IsPaid      bool
}

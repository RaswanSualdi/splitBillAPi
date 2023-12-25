package models

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	ID           int
	SplitBills   []SplitBill `gorm:"many2many:splitbill_customers;"`
	Name         string
	EMoney       int64
	Email        string
	PasswordHash string
}

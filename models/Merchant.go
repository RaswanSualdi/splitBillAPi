package models

import (
	"github.com/jinzhu/gorm"
)

type Merchant struct {
	gorm.Model
	ID           int
	Products     []Product
	MerchantName string
	EMoney       int64
}

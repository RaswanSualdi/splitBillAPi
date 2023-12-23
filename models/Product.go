package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	SplitBillID string
	MerchantID  string
	ProductName string
	Price       int64
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func (s *SplitBill) BeforeCreateProduct(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

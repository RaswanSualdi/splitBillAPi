package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Merchant struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	Products     []Product
	MerchantName string
	EMoney       int64
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

func (s *SplitBill) BeforeCreateMerchant(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

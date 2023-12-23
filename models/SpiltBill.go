package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type SplitBill struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Customers   []Customer
	Products    []Product
	Description string
	TotalAmount int64
	IsEnough    bool
	IsPaid      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *SplitBill) BeforeCreateSplitBill(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

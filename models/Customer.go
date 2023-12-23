package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Customer struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	SplitBillID  string
	Name         string
	EMoney       int64
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *SplitBill) BeforeCreateCustomer(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

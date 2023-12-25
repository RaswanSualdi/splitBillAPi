package usecase

import (
	"SplitBillApi/models"
	"SplitBillApi/payload/request"
	"SplitBillApi/repositories"
)

type UseCase interface {
	SplitBill(input request.SplitBillRequest) (models.SplitBill, error)
	// SplitBillPayment(input request.SplitBillPaymentReq) models.SplitBill
	GetSplitBillByCustomerID(id string) ([]models.SplitBill, error)
}

type usecase struct {
	repositories repositories.Repositories
}

func NewUseCase(repo repositories.Repositories) *usecase {
	return &usecase{repositories: repo}
}

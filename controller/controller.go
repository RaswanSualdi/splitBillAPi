package controller

import (
	"SplitBillApi/usecase"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	SplitBill(c *gin.Context)
}

type controller struct {
	usecase usecase.UseCase
}

func NewController(usecase usecase.UseCase) *controller {
	return &controller{usecase: usecase}
}

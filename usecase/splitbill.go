package usecase

import (
	"SplitBillApi/models"
	"SplitBillApi/payload/request"
)

func (u *usecase) SplitBill(input request.SplitBillRequest) (models.SplitBill, error) {
	var customers []models.Customer
	var isEnough bool

	for _, i := range input.To {
		customer, _ := u.repositories.GetCustomerById(i.CustomerID)
		customers = append(customers, customer)
	}
	var items []models.Product
	var price int64
	for _, j := range input.Item {
		item, _ := u.repositories.GetProductById(j.ProductID)
		price += item.Price
		items = append(items, item)
	}

	if price <= input.Nominal {
		isEnough = true
	} else {
		isEnough = false
	}
	splitBill := models.SplitBill{
		Customers:   customers,
		Products:    items,
		Description: input.Description,
		TotalAmount: input.Nominal,
		TotalPrice:  price,
		IsEnough:    isEnough,
		IsPaid:      false,
	}

	newSplitBill, err := u.repositories.SaveSplitBill(splitBill)
	if err != nil {
		return newSplitBill, err
	}

	return newSplitBill, nil

}

// func (u *usecase) SplitBillPayment(input request.SplitBillPaymentReq) (models.SplitBill, error) {
// 	u.repositories.
// }

func (u *usecase) GetSplitBillByCustomerID(id string) ([]models.SplitBill, error) {
	splitBills, err := u.repositories.GetSplitBillByCustomerID(id)
	if err != nil {
		return splitBills, err
	}

	return splitBills, nil
}

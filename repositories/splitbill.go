package repositories

import "SplitBillApi/models"

// import "SplitBillApi/models"

func (r *repositories) GetSplitBillByCustomerID(name string) ([]models.SplitBill, error) {
	var customer models.Customer
	err := r.db.Where("name = ?", name).Find(&customer).Error
	if err != nil {
		return customer.SplitBills, err
	}
	splitBills := customer.SplitBills
	return splitBills, nil
}

func (r *repositories) SaveSplitBill(splitBill models.SplitBill) (models.SplitBill, error) {
	err := r.db.Create(&splitBill).Error
	if err != nil {
		return splitBill, err
	}

	return splitBill, nil

}

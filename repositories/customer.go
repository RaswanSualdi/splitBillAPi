package repositories

import "SplitBillApi/models"

func (r *repositories) GetCustomerById(id int) (models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("ID = ?", id).Find(&customer).Error
	if err != nil {
		return customer, nil
	}
	return customer, nil
}

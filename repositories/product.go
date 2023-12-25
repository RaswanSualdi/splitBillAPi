package repositories

import "SplitBillApi/models"

func (r *repositories) GetProductById(id int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, "ID = ?", id).Error
	if err != nil {
		return product, err
	}

	return product, nil

}

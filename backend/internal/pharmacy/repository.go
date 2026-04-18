package pharmacy

import "meditrack/configs"

func FindAllMedicines() ([]Medicine, error) {
	var medicines []Medicine
	err := configs.DB.Find(&medicines).Error
	return medicines, err
}

func CreatePharmacyOrder(order *PharmacyOrder) error {
	return configs.DB.Create(order).Error
}

func FindAllPharmacyOrders() ([]PharmacyOrder, error) {
	var orders []PharmacyOrder
	err := configs.DB.Find(&orders).Error
	return orders, err
}

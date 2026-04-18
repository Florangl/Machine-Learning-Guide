package pharmacy

func GetMedicinesService() ([]Medicine, error) {
	return FindAllMedicines()
}

func CreatePharmacyOrderService(req CreatePharmacyOrderRequest) error {
	order := PharmacyOrder{
		PrescriptionID: req.PrescriptionID,
		MedicineName:   req.MedicineName,
		Quantity:       req.Quantity,
		Status:         "pending",
	}

	return CreatePharmacyOrder(&order)
}

func GetPharmacyOrdersService() ([]PharmacyOrder, error) {
	return FindAllPharmacyOrders()
}
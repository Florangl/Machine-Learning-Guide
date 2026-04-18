package payment

import "meditrack/configs"

func FindAllPayments() ([]Payment, error) {
	var payments []Payment
	err := configs.DB.Find(&payments).Error
	return payments, err
}

func FindPaymentByID(id int) (Payment, error) {
	var payment Payment
	err := configs.DB.First(&payment, id).Error
	return payment, err
}

func CreatePayment(payment *Payment) error {
	return configs.DB.Create(payment).Error
}
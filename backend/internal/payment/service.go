package payment

func GetPaymentsService() ([]Payment, error) {
	return FindAllPayments()
}

func GetPaymentByIDService(id int) (Payment, error) {
	return FindPaymentByID(id)
}

func CreatePaymentService(req CreatePaymentRequest) error {
	payment := Payment{
		AppointmentID: req.AppointmentID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: "paid",
	}

	return CreatePayment(&payment)
}
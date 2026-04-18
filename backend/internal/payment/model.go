package payment

type Payment struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	AppointmentID *uint   `json:"appointment_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	PaymentStatus string  `json:"payment_status"`
}

func (Payment) TableName() string {
	return "payments"
}

type CreatePaymentRequest struct {
	AppointmentID *uint   `json:"appointment_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}
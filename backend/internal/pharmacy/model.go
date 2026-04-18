package pharmacy

type Medicine struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}

func (Medicine) TableName() string {
	return "medicines"
}

type PharmacyOrder struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	PrescriptionID *uint `json:"prescription_id"`
	MedicineName string `json:"medicine_name"`
	Quantity     int    `json:"quantity"`
	Status       string `json:"status"`
}

func (PharmacyOrder) TableName() string {
	return "pharmacy_orders"
}

type CreatePharmacyOrderRequest struct {
	PrescriptionID *uint  `json:"prescription_id"`
	MedicineName   string `json:"medicine_name"`
	Quantity       int    `json:"quantity"`
}
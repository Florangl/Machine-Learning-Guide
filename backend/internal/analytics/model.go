package analytics

type PatientOutcomeReport struct {
	TotalMedicalRecords int64 `json:"total_medical_records"`
}

type DoctorPerformanceReport struct {
	DoctorID          uint  `json:"doctor_id"`
	TotalAppointments int64 `json:"total_appointments"`
}

type DrugUsageReport struct {
	MedicineName string `json:"medicine_name"`
	TotalOrders  int64  `json:"total_orders"`
}
package ehr

type MedicalRecord struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	PatientID   uint   `json:"patient_id"`
	DoctorID    uint   `json:"doctor_id"`
	VisitDate   string `json:"visit_date"`
	Diagnosis   string `json:"diagnosis"`
	Prescription string `json:"prescription"`
	LabResult   string `json:"lab_result"`
	Notes       string `json:"notes"`
}

func (MedicalRecord) TableName() string {
	return "medical_records"
}

type CreateMedicalRecordRequest struct {
	PatientID    uint   `json:"patient_id"`
	DoctorID     uint   `json:"doctor_id"`
	VisitDate    string `json:"visit_date"`
	Diagnosis    string `json:"diagnosis"`
	Prescription string `json:"prescription"`
	LabResult    string `json:"lab_result"`
	Notes        string `json:"notes"`
}
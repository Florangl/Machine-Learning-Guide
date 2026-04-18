package analytics

import "meditrack/configs"

func GetPatientOutcomeData() (PatientOutcomeReport, error) {
	var total int64
	err := configs.DB.Table("medical_records").Count(&total).Error
	return PatientOutcomeReport{
		TotalMedicalRecords: total,
	}, err
}

func GetDoctorPerformanceData() ([]DoctorPerformanceReport, error) {
	var reports []DoctorPerformanceReport

	err := configs.DB.
		Table("appointments").
		Select("doctor_id, COUNT(*) as total_appointments").
		Group("doctor_id").
		Scan(&reports).Error

	return reports, err
}

func GetDrugUsageData() ([]DrugUsageReport, error) {
	var reports []DrugUsageReport

	err := configs.DB.
		Table("pharmacy_orders").
		Select("medicine_name, COUNT(*) as total_orders").
		Group("medicine_name").
		Scan(&reports).Error

	return reports, err
}
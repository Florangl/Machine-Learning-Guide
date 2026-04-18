package analytics

func GetPatientOutcomesService() (PatientOutcomeReport, error) {
	return GetPatientOutcomeData()
}

func GetDoctorPerformanceService() ([]DoctorPerformanceReport, error) {
	return GetDoctorPerformanceData()
}

func GetDrugUsageService() ([]DrugUsageReport, error) {
	return GetDrugUsageData()
}
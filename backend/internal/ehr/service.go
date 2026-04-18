package ehr

func GetMedicalRecordsService() ([]MedicalRecord, error) {
	return FindAllMedicalRecords()
}

func CreateMedicalRecordService(req CreateMedicalRecordRequest) error {
	record := MedicalRecord{
		PatientID:    req.PatientID,
		DoctorID:     req.DoctorID,
		VisitDate:    req.VisitDate,
		Diagnosis:    req.Diagnosis,
		Prescription: req.Prescription,
		LabResult:    req.LabResult,
		Notes:        req.Notes,
	}

	return CreateMedicalRecord(&record)
}
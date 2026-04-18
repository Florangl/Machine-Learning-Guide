package ehr

import "meditrack/configs"

func FindAllMedicalRecords() ([]MedicalRecord, error) {
	var records []MedicalRecord
	err := configs.DB.Find(&records).Error
	return records, err
}

func CreateMedicalRecord(record *MedicalRecord) error {
	return configs.DB.Create(record).Error
}
package appointment

type Appointment struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	PatientID       uint   `json:"patient_id"`
	DoctorID        uint   `json:"doctor_id"`
	ScheduleID      uint   `json:"schedule_id"`
	AppointmentDate string `json:"appointment_date"`
	AppointmentTime string `json:"appointment_time"`
	Status          string `json:"status"`
	Complaint       string `json:"complaint"`
}

func (Appointment) TableName() string {
	return "appointments"
}

type CreateAppointmentRequest struct {
	PatientID       uint   `json:"patient_id"`
	DoctorID        uint   `json:"doctor_id"`
	ScheduleID      uint   `json:"schedule_id"`
	AppointmentDate string `json:"appointment_date"`
	AppointmentTime string `json:"appointment_time"`
	Complaint       string `json:"complaint"`
}

type RescheduleAppointmentRequest struct {
	AppointmentDate string `json:"appointment_date"`
	AppointmentTime string `json:"appointment_time"`
}
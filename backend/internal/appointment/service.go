package appointment

func GetAppointmentsService() ([]Appointment, error) {
	return FindAllAppointments()
}

func CreateAppointmentService(req CreateAppointmentRequest) error {
	appointment := Appointment{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		ScheduleID:      req.ScheduleID,
		AppointmentDate: req.AppointmentDate,
		AppointmentTime: req.AppointmentTime,
		Status:          "booked",
		Complaint:       req.Complaint,
	}

	return CreateAppointment(&appointment)
}

func RescheduleAppointmentService(id int, req RescheduleAppointmentRequest) error {
	return UpdateAppointmentSchedule(id, req.AppointmentDate, req.AppointmentTime)
}

func CancelAppointmentService(id int) error {
	return CancelAppointment(id)
}
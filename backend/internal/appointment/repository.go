package appointment

import "meditrack/configs"

func FindAllAppointments() ([]Appointment, error) {
	var appointments []Appointment
	err := configs.DB.Find(&appointments).Error
	return appointments, err
}

func CreateAppointment(appointment *Appointment) error {
	return configs.DB.Create(appointment).Error
}

func UpdateAppointmentSchedule(id int, appointmentDate string, appointmentTime string) error {
	return configs.DB.Model(&Appointment{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"appointment_date": appointmentDate,
			"appointment_time": appointmentTime,
			"status":           "rescheduled",
		}).Error
}

func CancelAppointment(id int) error {
	return configs.DB.Model(&Appointment{}).
		Where("id = ?", id).
		Update("status", "cancelled").Error
}
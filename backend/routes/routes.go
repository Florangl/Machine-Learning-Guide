package routes

import (
	"meditrack/internal/analytics"
	"meditrack/internal/appointment"
	"meditrack/internal/ehr"
	"meditrack/internal/payment"
	"meditrack/internal/pharmacy"
	"meditrack/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "MediTrack API is running",
			})
		})

		api.GET("/users/:id", user.GetUserByID)

		api.GET("/appointments", appointment.GetAppointments)
		api.POST("/appointments", appointment.BookAppointment)
		api.PUT("/appointments/:id/reschedule", appointment.RescheduleAppointment)
		api.PUT("/appointments/:id/cancel", appointment.CancelAppointmentHandler)

		api.GET("/ehr/records", ehr.GetMedicalRecords)
		api.POST("/ehr/records", ehr.CreateMedicalRecordHandler)

		api.GET("/pharmacy/medicines", pharmacy.GetMedicines)
		api.POST("/pharmacy/orders", pharmacy.CreatePharmacyOrderHandler)
		api.GET("/pharmacy/orders", pharmacy.GetPharmacyOrders)

		api.GET("/payments", payment.GetPayments)
		api.GET("/payments/:id", payment.GetPaymentByID)
		api.POST("/payments", payment.CreatePaymentHandler)

		api.GET("/analytics/patient-outcomes", analytics.GetPatientOutcomes)
		api.GET("/analytics/doctor-performance", analytics.GetDoctorPerformance)
		api.GET("/analytics/drug-usage", analytics.GetDrugUsage)

	}
}

package appointment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAppointments(c *gin.Context) {
	appointments, err := GetAppointmentsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch appointments",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "appointments fetched successfully",
		"data":    appointments,
	})
}

func BookAppointment(c *gin.Context) {
	var req CreateAppointmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := CreateAppointmentService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create appointment",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "appointment booked successfully",
	})
}

func RescheduleAppointment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid appointment id",
		})
		return
	}

	var req RescheduleAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err = RescheduleAppointmentService(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to reschedule appointment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "appointment rescheduled successfully",
	})
}

func CancelAppointmentHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid appointment id",
		})
		return
	}

	err = CancelAppointmentService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to cancel appointment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "appointment cancelled successfully",
	})
}
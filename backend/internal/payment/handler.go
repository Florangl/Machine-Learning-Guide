package payment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPayments(c *gin.Context) {
	payments, err := GetPaymentsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch payments",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "payments fetched successfully",
		"data":    payments,
	})
}

func GetPaymentByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payment id",
		})
		return
	}

	payment, err := GetPaymentByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "payment not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "payment fetched successfully",
		"data":    payment,
	})
}

func CreatePaymentHandler(c *gin.Context) {
	var req CreatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := CreatePaymentService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create payment",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "payment created successfully",
	})
}
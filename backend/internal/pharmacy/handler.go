package pharmacy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMedicines(c *gin.Context) {
	medicines, err := GetMedicinesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch medicines",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "medicines fetched successfully",
		"data":    medicines,
	})
}

func CreatePharmacyOrderHandler(c *gin.Context) {
	var req CreatePharmacyOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := CreatePharmacyOrderService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create pharmacy order",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "pharmacy order created successfully",
	})
}

func GetPharmacyOrders(c *gin.Context) {
	orders, err := GetPharmacyOrdersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch pharmacy orders",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pharmacy orders fetched successfully",
		"data":    orders,
	})
}
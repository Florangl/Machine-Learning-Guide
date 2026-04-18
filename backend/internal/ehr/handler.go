package ehr

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMedicalRecords(c *gin.Context) {
	records, err := GetMedicalRecordsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch medical records",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "medical records fetched successfully",
		"data":    records,
	})
}

func CreateMedicalRecordHandler(c *gin.Context) {
	var req CreateMedicalRecordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := CreateMedicalRecordService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create medical record",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "medical record created successfully",
	})
}
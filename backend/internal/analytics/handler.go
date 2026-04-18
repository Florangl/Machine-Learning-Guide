package analytics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPatientOutcomes(c *gin.Context) {
	report, err := GetPatientOutcomesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch patient outcomes",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "patient outcomes fetched successfully",
		"data":    report,
	})
}

func GetDoctorPerformance(c *gin.Context) {
	reports, err := GetDoctorPerformanceService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch doctor performance",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "doctor performance fetched successfully",
		"data":    reports,
	})
}

func GetDrugUsage(c *gin.Context) {
	reports, err := GetDrugUsageService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch drug usage",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "drug usage fetched successfully",
		"data":    reports,
	})
}
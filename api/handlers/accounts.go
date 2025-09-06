package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verma29897/bulksms/db"
	"github.com/verma29897/bulksms/models"
)

func ListAccounts(c *gin.Context) {
    if db.GormDB == nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{"error": "database not initialized"})
        return
    }
	var accounts []models.Account
	if err := db.GormDB.Limit(100).Order("waba_id ASC").Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

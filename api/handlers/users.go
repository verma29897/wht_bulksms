package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/verma29897/bulksms/db"
    "github.com/verma29897/bulksms/models"
)

func ListUsers(c *gin.Context) {
    if db.GormDB == nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{"error": "database not initialized"})
        return
    }
    var users []models.User
    if err := db.GormDB.Limit(50).Order("id DESC").Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    for i := range users {
        users[i].PasswordHash = "" // do not leak
    }
    c.JSON(http.StatusOK, gin.H{"users": users})
}



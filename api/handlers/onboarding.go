package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/verma29897/bulksms/models"
)

type storeOnboardingRequest struct {
    AccessToken string `json:"access_token"`
}

func StoreOnboarding(c *gin.Context) {
    var req storeOnboardingRequest
    if err := c.ShouldBindJSON(&req); err != nil || req.AccessToken == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "access_token is required"})
        return
    }

    account, err := models.FetchWABAData(req.AccessToken)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if uid, ok := c.Get("user_id"); ok {
        if id64, ok2 := uid.(int64); ok2 {
            account.UserID = &id64
        }
    }

    if err := models.StoreAccount(account); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Onboarding data stored", "waba_id": account.WABAID, "phone_number_id": account.PhoneNumberID})
}



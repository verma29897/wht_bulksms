package handlers

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/verma29897/bulksms/models"
)

func OnboardingCallback(c *gin.Context) {
    code := c.Query("auth_code")
    if code == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing auth_code"})
        return
    }

    token, err := exchangeAuthCode(code)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    account, err := models.FetchWABAData(token)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    err = models.StoreAccount(account)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Signup completed"})
}

func exchangeAuthCode(code string) (string, error) {
    appID := os.Getenv("APP_ID")
    appSecret := os.Getenv("APP_SECRET")
    redirectURI := os.Getenv("REDIRECT_URI")

    url := fmt.Sprintf("https://graph.facebook.com/v20.0/oauth/access_token?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s",
        appID, appSecret, redirectURI, code)

    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    var result map[string]interface{}
    _ = json.Unmarshal(body, &result)

    if resp.StatusCode != 200 {
        return "", fmt.Errorf("failed to exchange code: %s", body)
    }

    return result["access_token"].(string), nil
}

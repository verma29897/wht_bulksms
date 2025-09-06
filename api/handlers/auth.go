package handlers

import (
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "github.com/verma29897/bulksms/models"
)

type signupRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type loginRequest struct {
    Identifier string `json:"identifier"` // email or username
    Password   string `json:"password"`
}

func Signup(c *gin.Context) {
    var req signupRequest
    if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Username == "" || req.Password == "" || req.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
        return
    }

    email := strings.ToLower(strings.TrimSpace(req.Email))
    username := strings.ToLower(strings.TrimSpace(req.Username))
    existing, err := models.GetUserByEmail(email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if existing != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
        return
    }
    byUser, err := models.GetUserByUsername(username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if byUser != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Username already in use"})
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    _, err = models.CreateUser(req.Name, email, username, string(hash))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Account created"})
}

func Login(c *gin.Context) {
    var req loginRequest
    if err := c.ShouldBindJSON(&req); err != nil || req.Identifier == "" || req.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
        return
    }
    ident := strings.ToLower(strings.TrimSpace(req.Identifier))
    var user *models.User
    var err error
    if strings.Contains(ident, "@") {
        user, err = models.GetUserByEmail(ident)
    } else {
        user, err = models.GetUserByUsername(ident)
    }
    if err != nil || user == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "dev-secret" // for local use only
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "email": user.Email,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
    })
    s, err := token.SignedString([]byte(secret))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": s})
}

// Me returns the latest onboarded WABA and phone number for the logged-in user
func Me(c *gin.Context) {
    uidAny, ok := c.Get("user_id")
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
        return
    }
    uid, ok := uidAny.(int64)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
        return
    }
    acct, err := models.GetLatestAccountByUserID(uid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if acct == nil {
        c.JSON(http.StatusOK, gin.H{"waba_id": "", "phone_number_id": ""})
        return
    }
    c.JSON(http.StatusOK, gin.H{"waba_id": acct.WABAID, "phone_number_id": acct.PhoneNumberID})
}



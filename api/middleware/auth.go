package middleware

import (
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        header := c.GetHeader("Authorization")
        if header == "" || !strings.HasPrefix(strings.ToLower(header), "bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
            return
        }
        tokenStr := strings.TrimSpace(header[len("Bearer "):])
        secret := os.Getenv("JWT_SECRET")
        if secret == "" {
            secret = "dev-secret"
        }
        token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
            return
        }
        // Set user info into context
        if v, ok := claims["sub"].(float64); ok {
            c.Set("user_id", int64(v))
        }
        if v, ok := claims["email"].(string); ok {
            c.Set("user_email", v)
        }
        c.Next()
    }
}



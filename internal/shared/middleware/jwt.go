package middleware

import (
	"fmt"
	"net/http"
	"os"
	"reka-storage/internal/shared/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractTokenFromHeader(c)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		claims, err := validateToken(tokenString)
		if err != nil {
			if err.Error() == "JWT_SECRET is not configured" {
				response.Error(c, http.StatusInternalServerError, err.Error())
			} else {
				response.Error(c, http.StatusUnauthorized, "Invalid or expired token")
			}
			c.Abort()
			return
		}

		// Set userID to context
		if userID, ok := claims["user_id"].(string); ok {
			c.Set("userID", userID)
		}

		c.Next()
	}
}

func extractTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("Invalid authorization header format")
	}

	return parts[1], nil
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is not configured")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil
}

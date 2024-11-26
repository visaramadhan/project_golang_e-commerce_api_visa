package middleware

import (
	"fmt"
	"net/http"
	"strings"

	// "github.com/e-commerce-api/dto"
	// "github.com/e-commerce-api/helper"
	// "github.com/e-commerce-api/utils/security"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// type authHeader struct {
// 	AuthorizationHeader string `header:"Authorization"`
// }

// Middleware sederhana untuk otentikasi
var jwtSecret = []byte("your-secret-key")

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "Authorization header is missing"})
		c.Abort()
		return
	}

	// Extract Bearer token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "Invalid authorization header format"})
		c.Abort()
		return
	}

	tokenString := parts[1]

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "invalid to parse token"})
		c.Abort()
		return
	}

	// Validate the token
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "Invalid token"})
		c.Abort()
		return
	}

	// Retrieve claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "Invalid claims"})
		c.Abort()
		return
	}

	// Extract user_id
	userID := claims["user_id"].(string)
	c.Set("user_id", userID)

	c.Next()
}

// func ValidationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var input struct {
// 			Username string `json:"username" binding:"required"`
// 			Password string `json:"password" binding:"required"`
// 		}
// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			c.Abort()
// 		} else {
// 			c.Set("user", input)
// 			c.Next()
// 		}
// 	}
// }

// Middleware untuk memeriksa otorisasi admin
// func AdminAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		role := c.GetString("role")

// 		if role != "admin" {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

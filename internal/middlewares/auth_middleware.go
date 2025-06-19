package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	RoleID int64 `json:"role_id"`
	jwt.RegisteredClaims
}

// func AuthMiddleware(secretKey string, allowedRoles ...int64) gin.HandlerFunc{

// }
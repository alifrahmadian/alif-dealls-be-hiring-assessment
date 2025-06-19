package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/handlers/responses"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	RoleID int64 `json:"role_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware(secretKey string, allowedRoles ...int64) gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.ErrorResponse(c, http.StatusUnauthorized, e.ErrNoAuthorizationHeader.Error())
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString != authHeader {
			responses.ErrorResponse(c, http.StatusUnauthorized, e.ErrInvalidTokenFormat.Error())
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			responses.ErrorResponse(c, http.StatusUnauthorized, e.ErrTokenExpired.Error())
			c.Abort()
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			responses.ErrorResponse(c, http.StatusUnauthorized, e.ErrTokenExpired.Error())
			c.Abort()
			return
		}

		roleAllowed := false
		for _, role := range allowedRoles {
			if claims.RoleID == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			responses.ErrorResponse(c, http.StatusForbidden, e.ErrForbidden.Error())
			c.Abort()
			return
		}

		c.Set("user_id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("role_id", claims.RoleID)

		c.Next()
	}
}
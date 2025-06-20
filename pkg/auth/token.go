package auth

import (
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	RoleID int64 `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User, secretKey string, ttl int) (string, error) {
	jwtSecret := []byte(secretKey)

	claims := Claims {
		ID: user.ID,
		Username: user.Username,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(ttl))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
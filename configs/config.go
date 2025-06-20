package configs

import (
	"database/sql"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/handlers"
)

type Config struct {
	DB *sql.DB
	Auth *AuthConfig
	Handler *Handler
	Env string
}

type Handler struct {
	AuthHandler *handlers.AuthHandler
	AttendancePeriodHandler *handlers.AttendancePeriodHandler
}

type AuthConfig struct {
	TTL int
	SecretKey string
}
package routes

import (
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/configs"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/constants"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(secretKey string, router *gin.Engine, handlers *configs.Handler){
	publicRoutes := router.Group("")
	{
		publicRoutes.POST("/login", handlers.AuthHandler.Login)
	}

	adminRoutes := router.Group("")
	adminRoutes.Use(middlewares.AuthMiddleware(secretKey, constants.RoleAdmin)) 
	{
		adminRoutes.POST("/admin/attendance_periods", handlers.AttendancePeriodHandler.CreateAttendancePeriod)
	}
	


}
package routes

import (
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/configs"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/constants"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(secretKey string, router *gin.Engine, handlers *configs.Handler) {
	publicRoutes := router.Group("")
	{
		publicRoutes.POST("/login", handlers.AuthHandler.Login)
	}

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middlewares.AuthMiddleware(secretKey, constants.RoleAdmin))
	{
		adminRoutes.POST("/attendance_periods", handlers.AttendancePeriodHandler.CreateAttendancePeriod)
		adminRoutes.POST("/payrolls", handlers.PayrollHandler.CreatePayroll)
	}

	employeeRoutes := router.Group("")
	employeeRoutes.Use(middlewares.AuthMiddleware(secretKey, constants.RoleEmployee))
	{
		employeeRoutes.POST("/attendances", handlers.AttendanceHandler.CreateAttendance)
		employeeRoutes.POST("/overtimes", handlers.OvertimeHandler.CreateOvertime)
		employeeRoutes.POST("/reimbursement", handlers.ReimbursementHandler.CreateReimbursement)
		employeeRoutes.GET("/payrolls/generate_payslip", handlers.PayrollHandler.GeneratePayslip)
	}
}

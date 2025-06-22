package main

import (
	"fmt"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/configs"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/db"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/handlers"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/repositories"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/routes"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/services"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *configs.Config
}

func LoadConfig() (*configs.Config, error) {
	err := configs.LoadGoDotEnv()
	if err != nil {
		return nil, err
	}

	dbConfig := configs.LoadDBConfig()
	env := configs.LoadEnv()
	authConfig := configs.LoadAuthConfig()

	db, err := db.Connect(*dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	userRepo := repositories.NewUserRepository(db)
	attendancePeriodRepo := repositories.NewAttendancePeriodRepository(db)
	attendanceRepo := repositories.NewAttendanceRepository(db)
	overtimeRepo := repositories.NewOvertimeRepository(db)
	reimbursementRepo := repositories.NewReimbursementRepository(db)
	payrollRepo := repositories.NewPayrollRepository(db)

	authService := services.NewAuthService(userRepo)
	attendancePeriodService := services.NewAttendancePeriodService(attendancePeriodRepo)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	overtimeService := services.NewOvertimeService(overtimeRepo, attendanceRepo)
	reimbursementService := services.NewReimbursementService(reimbursementRepo)
	payrollService := services.NewPayrollService(
		payrollRepo,
		userRepo,
		attendancePeriodRepo,
		attendanceRepo,
		overtimeRepo,
		reimbursementRepo,
	)

	authHandler := handlers.NewAuthHandler(&authService, authConfig.SecretKey, authConfig.TTL)
	attendancePeriodHandler := handlers.NewAttendancePeriodHandler(&attendancePeriodService)
	attendanceHandler := handlers.NewAttendanceHandler(&attendanceService)
	overtimeHandler := handlers.NewOvertimeHandler(&overtimeService)
	reimbursementHandler := handlers.NewReimbursementHandler(&reimbursementService)
	payrollHandler := handlers.NewPayrollHandler(&payrollService)

	return &configs.Config{
		DB:   db,
		Env:  env,
		Auth: authConfig,
		Handler: &configs.Handler{
			AuthHandler:             authHandler,
			AttendancePeriodHandler: attendancePeriodHandler,
			AttendanceHandler:       attendanceHandler,
			OvertimeHandler:         overtimeHandler,
			ReimbursementHandler:    reimbursementHandler,
			PayrollHandler:          payrollHandler,
		},
	}, nil
}

func NewApp() *App {
	cfg, err := LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(cfg.Auth.SecretKey, router, cfg.Handler)

	return &App{
		Router: router,
		Config: cfg,
	}
}

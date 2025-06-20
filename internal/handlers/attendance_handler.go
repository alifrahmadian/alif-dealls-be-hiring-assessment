package handlers

import (
	"net/http"
	"time"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/handlers/dtos"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/handlers/responses"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/services"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/messages"
	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	AttendanceService services.AttendanceService
}

func NewAttendanceHandler(attendanceService *services.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		AttendanceService: *attendanceService,
	}
}

func (h *AttendanceHandler) CreateAttendance(c *gin.Context) {
	userID := c.GetInt64("user_id")

	checkInDate := time.Now()
	if isWeekend(checkInDate) {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrIsWeekend.Error())
		return
	}
	

	attendance := &models.Attendance{
		UserID: userID,
		Date: checkInDate,
		CreatedBy: userID,
		UpdatedBy: userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newAttendance, err := h.AttendanceService.CreateAttendance(attendance)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.CreateAttendanceResponse{
		ID: newAttendance.ID,
		UserID: newAttendance.UserID,
		Date: newAttendance.Date.Format(time.RFC3339),
	}

	responses.SuccessResponse(c, messages.RspCreateAttendanceSuccess, resp)
}
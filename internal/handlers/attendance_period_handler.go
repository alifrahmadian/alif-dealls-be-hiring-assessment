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
	"github.com/go-playground/validator/v10"
)

type AttendancePeriodHandler struct {
	AttendancePeriodService services.AttendancePeriodService
}

func NewAttendancePeriodHandler(attendancePeriodService *services.AttendancePeriodService) *AttendancePeriodHandler {
	return &AttendancePeriodHandler{
		AttendancePeriodService: *attendancePeriodService,
	}
}

func (h *AttendancePeriodHandler) CreateAttendancePeriod(c *gin.Context) {
	var req dtos.CreateAttendancePeriodRequest

	userID := c.GetInt64("user_id")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "StartDate":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrStartDateRequired.Error())
				return
			case "EndDate":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrEndDateRequired.Error())
				return
			}
		}
	}

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, req.StartDate)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrInvalidStartDate.Error())
		return
	}

	endDate, err := time.Parse(layout, req.EndDate)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrInvalidEndDate.Error())
		return
	}

	attendancePeriod := &models.AttendancePeriod{
		StartDate: startDate,
		EndDate: endDate,
		CreatedBy: userID,
		UpdatedBy: userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newAttendancePeriod, err := h.AttendancePeriodService.CreateAttendancePeriod(attendancePeriod)
	if err != nil{
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.CreateAttendancePeriodResponse{
		ID: newAttendancePeriod.ID,
		StartDate: newAttendancePeriod.StartDate.Format("2006-01-02"),
		EndDate: newAttendancePeriod.EndDate.Format("2006-01-02"),
		CreatedBy: newAttendancePeriod.CreatedBy,
		UpdatedBy: newAttendancePeriod.UpdatedBy,
		CreatedAt: newAttendancePeriod.CreatedAt.Format(time.RFC3339),
		UpdatedAt: newAttendancePeriod.UpdatedAt.Format(time.RFC3339),
	}

	responses.SuccessResponse(c, messages.RspCreateAttendencePeriodSuccess, resp)
}
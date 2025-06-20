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

type OvertimeHandler struct {
	OvertimeService services.OvertimeService
}

func NewOvertimeHandler(overtimeService *services.OvertimeService) *OvertimeHandler {
	return &OvertimeHandler{
		OvertimeService: *overtimeService,
	}
}

func (h *OvertimeHandler) CreateOvertime(c *gin.Context) {
	var req dtos.CreateOvertimeRequest

	userID := c.GetInt64("user_id")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Date":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrOvertimeDateIsRequired.Error())
				return
			case "OvertimeHours":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrOvertimeHoursRequired.Error())
				return
			}
		}
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, req.Date)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrInvalidOvertimeDate.Error())
		return
	}

	overtime := &models.Overtime{
		UserID: userID,
		Date: date,
		OvertimeHours: req.OvertimeHours,
		CreatedBy: userID,
		UpdatedBy: userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newOvertime, err := h.OvertimeService.CreateOvertime(overtime)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err == e.ErrEmployeeHasNotWorkedToday {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrEmployeeHasNotWorkedToday.Error())
		return
	}

	if err == e.ErrOvertimeMoreThanThreeHours {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrOvertimeDateIsRequired.Error())
		return
	}

	if err == e.ErrOvertimeIsTakenToday {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrOvertimeIsTakenToday.Error())
		return
	}

	resp := &dtos.CreateOvertimeResponse{
		ID: newOvertime.ID,
		UserID: newOvertime.UserID,
		Date: newOvertime.Date.Format("2006-01-02"),
		OvertimeHours: newOvertime.OvertimeHours,
	}

	responses.SuccessResponse(c, messages.RspCreateOvertimeSuccess, resp)
}
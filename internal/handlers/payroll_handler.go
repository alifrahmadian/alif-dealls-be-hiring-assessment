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

type PayrollHandler struct {
	PayrollService services.PayrollService
}

func NewPayrollHandler(payrollService *services.PayrollService) *PayrollHandler {
	return &PayrollHandler{
		PayrollService: *payrollService,
	}
}

func (h *PayrollHandler) CreatePayroll(c *gin.Context) {
	var req dtos.CreatePayrollRequest

	userID := c.GetInt64("user_id")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "UserID":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrPayrollUserIDRequired.Error())
				return
			case "AttendancePeriodID":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrPayrollAttendancePeriodIDRequired.Error())
				return
			}
		}
	}

	payroll := &models.Payroll{
		UserID:             req.UserID,
		AttendancePeriodID: req.AttendancePeriodID,
		CreatedBy:          userID,
		UpdatedBy:          userID,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	newPayroll, err := h.PayrollService.CreatePayroll(payroll)
	if err != nil {
		if err == e.ErrUserNotFound {
			responses.ErrorResponse(c, http.StatusNotFound, e.ErrUserNotFound.Error())
			return
		}

		if err == e.ErrAttendancePeriodNotFound {
			responses.ErrorResponse(c, http.StatusNotFound, e.ErrAttendancePeriodNotFound.Error())
			return
		}

		if err == e.ErrPayrollHasBeenProcessed {
			responses.ErrorResponse(c, http.StatusConflict, e.ErrPayrollHasBeenProcessed.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.CreatePayrollResponse{
		ID:                  newPayroll.ID,
		UserID:              newPayroll.UserID,
		AttendancePeriodID:  newPayroll.AttendancePeriodID,
		BaseSalary:          newPayroll.BaseSalary,
		AttendanceDays:      newPayroll.AttendanceDays,
		AttendanceAmount:    newPayroll.AttendanceAmount,
		OvertimeHours:       newPayroll.OvertimeHours,
		OvertimeAmount:      newPayroll.OvertimeAmount,
		ReimbursementAmount: newPayroll.ReimbursementAmount,
		TotalTakeHomePay:    newPayroll.TotalTakeHomePay,
		CreatedBy:           newPayroll.CreatedBy,
		UpdatedBy:           newPayroll.UpdatedBy,
		CreatedAt:           newPayroll.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           newPayroll.UpdatedAt.Format(time.RFC3339),
	}

	responses.SuccessResponse(c, messages.RspCreatePayrollSuccess, resp)

}

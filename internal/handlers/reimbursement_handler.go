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

type ReimbursementHandler struct {
	ReimbursementService services.ReimbursementService
}

func NewReimbursementHandler(reimbursementService *services.ReimbursementService) *ReimbursementHandler {
	return &ReimbursementHandler{
		ReimbursementService: *reimbursementService,
	}
}

func (h *ReimbursementHandler) CreateReimbursement(c *gin.Context) {
	var req dtos.CreateReimbursementRequest

	userID := c.GetInt64("user_id")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors){
			switch err.Field() {
			case "ReimbursementAmount":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrReimbursementAmountIsRequired.Error())
				return
			case "Date":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrReimbursementDateIsRequired.Error())
				return
			case "Description":
				responses.ErrorResponse(c, http.StatusBadRequest, e.ErrReimbursementDescriptionIsRequired.Error())
				return
			}
		}
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, req.Date)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, e.ErrInvalidReimbursementDate.Error())
		return
	}

	reimbursement := &models.Reimbursement{
		UserID: userID,
		ReimbursementAmount: req.ReimbursementAmount,
		Date: date,
		Description: req.Description,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	newReimbursement, err := h.ReimbursementService.CreateReimbursement(reimbursement)
	if err != nil {
		if err == e.ErrReimbursementDateIsAfterToday {
			responses.ErrorResponse(c, http.StatusBadRequest, e.ErrReimbursementDateIsAfterToday.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.CreateReimbursementResponse{
		ID: newReimbursement.ID,
		UserID: newReimbursement.UserID,
		ReimbursementAmount: newReimbursement.ReimbursementAmount,
		Date: newReimbursement.Date.Format("2006-01-02"),
		Description: newReimbursement.Description,
		CreatedBy: newReimbursement.CreatedBy,
		UpdatedBy: newReimbursement.UpdatedBy,
		CreatedAt: newReimbursement.CreatedAt.Format(time.RFC3339),
		UpdatedAt: newReimbursement.UpdatedAt.Format(time.RFC3339),
	}

	responses.SuccessResponse(c, messages.RspCreateReimbursementSuccess, resp)
}
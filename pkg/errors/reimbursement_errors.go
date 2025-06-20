package errors

import "errors"

var (
	ErrReimbursementDateIsAfterToday = errors.New("cannot propose reimbursement after today")
	ErrReimbursementAmountIsRequired = errors.New("reimbursement amount is required")
	ErrReimbursementDateIsRequired = errors.New("reimbursement date is required")
	ErrReimbursementDescriptionIsRequired = errors.New("reimbursement description is required")
	ErrInvalidReimbursementDate = errors.New("invalid reimbursement date")
)
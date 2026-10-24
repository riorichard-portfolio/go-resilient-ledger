package entity

import "go-resilient-ledger/internal/pkg/apperr"

var(
	InvalidBalanceErr = apperr.BadRequestErr("invalid balance number")
	InvalidAmountErr = apperr.BadRequestErr("invalid amount number")
	InsufficientBalanceErr = apperr.BadRequestErr("insufficient balance")
)
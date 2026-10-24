package entity

import "go-resilient-ledger/internal/pkg/apperr"

var(
	InvalidAmountErr = apperr.BadRequestErr("invalid amount number")
)
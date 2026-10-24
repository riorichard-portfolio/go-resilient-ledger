package entity

import "go-resilient-ledger/internal/pkg/err/bizerr"

var(
	InvalidBalanceErr = bizerr.BadRequestErr("invalid balance number")
	InvalidAmountErr = bizerr.BadRequestErr("invalid amount number")
	InsufficientBalanceErr = bizerr.BadRequestErr("insufficient balance")
)
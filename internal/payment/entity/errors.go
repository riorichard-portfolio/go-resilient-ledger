package entity

import "go-resilient-ledger/internal/pkg/err/bizerr"

var(
	InvalidAmountErr = bizerr.BadRequestErr("invalid amount number")
)
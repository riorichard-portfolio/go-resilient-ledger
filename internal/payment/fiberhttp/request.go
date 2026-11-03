package fiberhttp

import "go-resilient-ledger/internal/pkg/listf"

type CreatePayment struct {
	AccountID string `json:"account_id" validate:"required"`
	Amount    int    `json:"amount" validate:"required"`
}

type FindPayment struct {
	AccountID string `json:"account_id" validate:"required"`
	listf.FilterJson
}

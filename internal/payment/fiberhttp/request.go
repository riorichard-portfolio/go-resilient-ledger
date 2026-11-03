package fiberhttp

import "go-resilient-ledger/internal/pkg/listf"

type CreatePaymentReq struct {
	AccountID string `json:"account_id" validate:"required"`
	Amount    int    `json:"amount" validate:"required"`
}

type FindPaymentsReq struct {
	AccountID string `json:"account_id" validate:"required"`
	listf.FilterJson
}

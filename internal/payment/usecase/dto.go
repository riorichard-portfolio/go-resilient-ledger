package usecase

import "go-resilient-ledger/internal/pkg/listf"

type PaymentMapVal struct {
	ID        string
	AccountID string
	Amount    int
}

type BulkFindInput struct {
	ID    string
	ReqID string
	listf.Filter
}

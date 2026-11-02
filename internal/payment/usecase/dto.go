package usecase

import (
	"time"

	"go-resilient-ledger/internal/pkg/listf"
)

type PaymentMapVal struct {
	ID        string
	AccountID string
	Amount    int
	CreatedAt time.Time
}

type BulkFindInput struct {
	AccountID string
	ReqID     string
	listf.Filter
}

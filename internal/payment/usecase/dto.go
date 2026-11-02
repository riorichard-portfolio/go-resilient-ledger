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
	QueryID     string
	listf.Filter
}

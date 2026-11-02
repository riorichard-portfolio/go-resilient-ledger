package findpaym

import (
	"time"

	"go-resilient-ledger/internal/pkg/listf"
)

type Input struct {
	AccountID string
	ReqID     string
	listf.Filter
}

type Payment struct {
	ID        string
	AccountID string
	Amount    int
	CreatedAt time.Time
}

type Output struct {
	Err      error
	Payments []Payment
}

type WQdata struct {
	Input
	Reply chan<- Output
}

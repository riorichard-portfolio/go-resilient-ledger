package findpaym

import "go-resilient-ledger/internal/pkg/listf"

type Input struct {
	ID    string
	ReqID string
	listf.Filter
}

type Payment struct {
	ID        string
	AccountID string
	Amount    int
}

type Output struct {
	Err      error
	Payments []Payment
}

type WQdata struct {
	Input
	Reply chan<- Output
}

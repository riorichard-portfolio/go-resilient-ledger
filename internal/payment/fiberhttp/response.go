package fiberhttp

import "time"

type Payment struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

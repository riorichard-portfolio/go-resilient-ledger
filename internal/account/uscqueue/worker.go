package uscqueue

import (
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
)

type Worker struct {
	createAccountWQ createacc.CreateAccountWQ
}
package uscqueue

import (
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
)

type Worker struct {
	createAccountWQ createacc.CreateAccountWQ
	findAccountWQ findacc.FindAccountWQ
}
package uscqueue

import (
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
)

type Producer struct {
	createAccountPQ createacc.CreateAccountPQ
}

func (p *Producer) AddCreateAccountQ(input createacc.Input) error {
	replyChan := make(chan error)
	p.createAccountPQ.WorkQueue() <- createacc.WorkQueue{
		Input: input,
		Reply: replyChan,
	}
	err := <-replyChan
	return err
}

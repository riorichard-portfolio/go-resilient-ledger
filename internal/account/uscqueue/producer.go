package uscqueue

import (
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
)

type Producer struct {
	createAccountPQ createacc.CreateAccountPQ
	findAccountPQ   findacc.FindAccountPQ
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

func (p *Producer) AddFindAccountQ(input findacc.Input) (findacc.Account, error) {
	replyChan := make(chan findacc.Output)
	p.findAccountPQ.WorkQueue() <- findacc.WorkQueue{
		Input: input,
		Reply: replyChan,
	}
	output := <-replyChan
	if output.Err != nil {
		return findacc.EmptyAccount, output.Err
	}
	return output.Account, nil
}

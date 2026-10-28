package uscqueue

import (
	"context"

	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
	"go-resilient-ledger/internal/pkg/apperr"
)

type Producer struct {
	createAccountPQ *createacc.CreateAccountPQ
	findAccountPQ   *findacc.FindAccountPQ
}

func (p *Producer) AddCreateAccountQ(ctx context.Context, input createacc.Input) error {
	replyChan := make(chan error, 1)
	select {
	case <-ctx.Done():
		return ctx.Err()

	case p.createAccountPQ.WorkQueue() <- createacc.WQdata{
		Input: input,
		Reply: replyChan,
	}:
		select {
		case err := <-replyChan:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	default:
		return apperr.TooManyRequestErr()
	}
}

func (p *Producer) AddFindAccountQ(ctx context.Context, input findacc.Input) (findacc.Account, error) {
	replyChan := make(chan findacc.Output, 1)
	select {
	case <-ctx.Done():
		return findacc.EmptyAccount, ctx.Err()

	case p.findAccountPQ.WorkQueue() <- findacc.WQdata{
		Input: input,
		Reply: replyChan,
	}:
		select {
		case output := <-replyChan:
			if output.Err != nil {
				return findacc.EmptyAccount, output.Err
			}
			return output.Account, nil
		case <-ctx.Done():
			return findacc.EmptyAccount, ctx.Err()
		}
	default:
		return findacc.EmptyAccount, apperr.TooManyRequestErr()
	}

}

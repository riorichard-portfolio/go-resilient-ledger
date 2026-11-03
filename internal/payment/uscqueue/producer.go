package uscqueue

import (
	"context"

	"go-resilient-ledger/internal/payment/uscqueue/queuechan/createpaym"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/findpaym"
	"go-resilient-ledger/internal/pkg/apperr"
)

type Producer struct {
	createPaymentPQ *createpaym.CreatePaymentPQ
	findPaymentPQ   *findpaym.FindPaymentPQ
}

func (p *Producer) AddCreatePaymentQ(ctx context.Context, input createpaym.Input) error {
	replyChan := make(chan error, 1)
	select {
	case <-ctx.Done():
		return ctx.Err()

	case p.createPaymentPQ.WorkQueue() <- createpaym.WQdata{
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

func (p *Producer) AddFindPaymentQ(ctx context.Context, input findpaym.Input) ([]findpaym.Payment, error) {
	replyChan := make(chan findpaym.Output, 1)
	select {
	case <-ctx.Done():
		return []findpaym.Payment{}, ctx.Err()

	case p.findPaymentPQ.WorkQueue() <- findpaym.WQdata{
		Input: input,
		Reply: replyChan,
	}:
		select {
		case output := <-replyChan:
			if output.Err != nil {
				return []findpaym.Payment{}, output.Err
			}
			return output.Payments, nil
		case <-ctx.Done():
			return []findpaym.Payment{}, ctx.Err()
		}
	default:
		return []findpaym.Payment{}, apperr.TooManyRequestErr()
	}

}

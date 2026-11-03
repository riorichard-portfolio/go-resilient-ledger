package uscqueue

import (
	"context"
	"log"

	"go-resilient-ledger/internal/payment/uscqueue/queuechan/createpaym"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/findpaym"
	"go-resilient-ledger/internal/payment/usecase"
	"go-resilient-ledger/internal/pkg/timer"
)

type Worker struct {
	createPaymentWQ *createpaym.CreatePaymentWQ
	findPaymentWQ   *findpaym.FindPaymentWQ
	writeBatch      int
	readBatch       int
	waitTimeoutMs   int
	usc             *usecase.Usecase
}

func (w *Worker) runCreatePaymentWorker(ctx context.Context) {
	batchWQData := make([]createpaym.WQdata, 0, w.writeBatch)
	t := timer.NewTimer(w.waitTimeoutMs)
	for {
		select {
		case <-ctx.Done():
			if len(batchWQData) > 0 {
				// usecase dipnggil
				w.usc.BulkCreatePayment(batchWQData)
			}
			return
		case data, ok := <-w.createPaymentWQ.WorkQueue():
			if !ok {
				if len(batchWQData) > 0 {
					// usecase dipnggil
					w.usc.BulkCreatePayment(batchWQData)
				}
				return
			}

			batchWQData = append(batchWQData, data)

			if len(batchWQData) == 1 {
				t.Start()
			}

			if len(batchWQData) >= w.writeBatch {
				t.Stop()
				// usecase dipnggil
				w.usc.BulkCreatePayment(batchWQData)
				batchWQData = batchWQData[:0]
			}
		case <-t.Timeout():
			if len(batchWQData) > 0 {
				// usecase dipnggil
				w.usc.BulkCreatePayment(batchWQData)
				batchWQData = batchWQData[:0]
			}
		}
	}
}

func (w *Worker) runFindPaymentWorker(ctx context.Context) {
	batchWQData := make([]findpaym.WQdata, 0, w.readBatch)
	t := timer.NewTimer(w.waitTimeoutMs)
	for {
		select {
		case <-ctx.Done():
			if len(batchWQData) > 0 {
				// usecase dipnggil
				w.usc.BulkFindPayments(batchWQData)
			}
			return
		case data, ok := <-w.findPaymentWQ.WorkQueue():
			if !ok {
				if len(batchWQData) > 0 {
					// usecase dipnggil
					w.usc.BulkFindPayments(batchWQData)
				}
				return
			}

			batchWQData = append(batchWQData, data)

			if len(batchWQData) == 1 {
				t.Start()
			}

			if len(batchWQData) >= w.readBatch {
				t.Stop()
				// usecase dipnggil
				w.usc.BulkFindPayments(batchWQData)
				batchWQData = batchWQData[:0]
			}
		case <-t.Timeout():
			if len(batchWQData) > 0 {
				// usecase dipnggil
				w.usc.BulkFindPayments(batchWQData)
				batchWQData = batchWQData[:0]
			}
		}
	}
}

func (w *Worker) StartWorker(ctx context.Context) {
	go w.runCreatePaymentWorker(ctx)
	go w.runFindPaymentWorker(ctx)
	log.Println("payment worker running")
}

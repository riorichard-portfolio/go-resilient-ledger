package uscqueue

import (
	"context"

	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
	"go-resilient-ledger/internal/pkg/timer"
)

type Worker struct {
	createAccountWQ createacc.CreateAccountWQ
	findAccountWQ   findacc.FindAccountWQ
	writeBatch      int
	readBatch       int
	waitTimeoutMs   int
}

func (w *Worker) runCreateAccountWorker(ctx context.Context) {
	batchWQData := make([]createacc.WQdata, 0, w.writeBatch)
	t := timer.NewTimer(w.waitTimeoutMs)
	for {
		select {
		case <-ctx.Done():
			if len(batchWQData) > 0 {
				// usecase dipnggil
			}
			return
		case data, ok := <-w.createAccountWQ.WorkQueue():
			if !ok {
				return
			}

			batchWQData = append(batchWQData, data)

			if len(batchWQData) == 1 {
				t.Start()
			}

			if len(batchWQData) >= w.writeBatch {
				t.Stop()
				// usecase dipnggil
				batchWQData = batchWQData[:0]
			}
		case <-t.Timeout():
			if len(batchWQData) > 0 {
				// usecase dipnggil
				batchWQData = batchWQData[:0]
			}
		}
	}
}

package uscqueue

import (
	"time"

	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
)

type Worker struct {
	createAccountWQ createacc.CreateAccountWQ
	findAccountWQ   findacc.FindAccountWQ
	writeBatch      int
	readBatch       int
	waitTimeout     time.Duration
}

func (w *Worker) runCreateAccountWorker() {
	batchWQData := make([]createacc.WQdata, 0, w.writeBatch)
	for {
		// select{
		// 	case w.createAccountWQ.WorkQueue()
		// }
	}
}

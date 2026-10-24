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
	batch
	for {
		// select{
		// 	case w.createAccountWQ.WorkQueue()
		// }
	}
}

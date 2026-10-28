package uscqueue

import (
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
)

func NewQueueSys(
	config Config,
) (*Worker, *Producer) {
	createAccWQChan := make(chan createacc.WQdata, config.CreateAccQSize)
	findAccWQChan := make(chan findacc.WQdata, config.FindAccQSize)
	createAccWQ, createAccPQ := createacc.NewCreateAccountWQ(createAccWQChan)
	findAccWQ, findAccPQ := findacc.NewFindAccountWQ(findAccWQChan)
	worker := Worker{
		createAccountWQ: createAccWQ,
		findAccountWQ:   findAccWQ,
		writeBatch:      config.WriteBatchSize,
		readBatch:       config.ReadBatchSize,
		waitTimeoutMs:   config.WaitTimeoutMs,
	}
	producer := Producer{
		createAccountPQ: createAccPQ,
		findAccountPQ:   findAccPQ,
	}
	return &worker, &producer
}

package uscqueue

import (
	"go-resilient-ledger/internal/account/usecase"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
)

func NewQueueSys(
	config Config,
	usc *usecase.Usecase,
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
		usc: usc,
	}
	producer := Producer{
		createAccountPQ: createAccPQ,
		findAccountPQ:   findAccPQ,
	}
	return &worker, &producer
}

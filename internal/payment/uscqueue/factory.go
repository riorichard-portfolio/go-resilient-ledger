package uscqueue

import (
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/createpaym"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/findpaym"
	"go-resilient-ledger/internal/payment/usecase"
)

func NewQueueSys(
	config Config,
	usc *usecase.Usecase,
) (*Worker, *Producer) {
	createPaymWQChan := make(chan createpaym.WQdata, config.CreatePaymQSize)
	findPaymWQChan := make(chan findpaym.WQdata, config.FindPaymQSize)
	createPaymWQ, createPaymPQ := createpaym.NewCreatePaymentWQ(createPaymWQChan)
	findPaymWQ, findPaymPQ := findpaym.NewFindPaymentWQ(findPaymWQChan)
	worker := Worker{
		createPaymentWQ: createPaymWQ,
		findPaymentWQ:   findPaymWQ,
		writeBatch:      config.WriteBatchSize,
		readBatch:       config.ReadBatchSize,
		waitTimeoutMs:   config.WaitTimeoutMs,
		usc:             usc,
	}
	producer := Producer{
		createPaymentPQ: createPaymPQ,
		findPaymentPQ:   findPaymPQ,
	}
	return &worker, &producer
}

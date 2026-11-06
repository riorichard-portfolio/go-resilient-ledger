package payment

import (
	"go-resilient-ledger/internal/payment/gormrepo"
	"go-resilient-ledger/internal/payment/gormtx"
	"go-resilient-ledger/internal/payment/uscqueue"
	"go-resilient-ledger/internal/payment/usecase"

	accsvc "go-resilient-ledger/internal/account/service"

	"gorm.io/gorm"
)

type Container struct {
	usc *usecase.Usecase

	QWorker   *uscqueue.Worker
	QProducer *uscqueue.Producer
}

func NewContainer(
	db *gorm.DB,
	accSvc *accsvc.Service,
) *Container {
	uscRepo := gormrepo.NewRepo(db)
	txManager := gormtx.NewGormTxManager(db)
	usc := usecase.NewUsecase(uscRepo, accSvc, txManager)
	w, p := uscqueue.NewQueueSys(uscqueue.Config{
		CreatePaymQSize: 10000,
		FindPaymQSize:   3000,
		ReadBatchSize:   100,
		WriteBatchSize:  5000,
		WaitTimeoutMs:   10,
	}, usc)
	return &Container{
		QWorker:   w,
		QProducer: p,
	}
}

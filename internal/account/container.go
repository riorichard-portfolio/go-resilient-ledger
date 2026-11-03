package account

import (
	"go-resilient-ledger/internal/account/gormrepo"
	"go-resilient-ledger/internal/account/service"
	"go-resilient-ledger/internal/account/uscqueue"
	"go-resilient-ledger/internal/account/usecase"

	"gorm.io/gorm"
)

type Container struct {
	usc *usecase.Usecase

	QWorker   *uscqueue.Worker
	QProducer *uscqueue.Producer
	Service   *service.Service
}

func NewContainer(
	db *gorm.DB,
) *Container {
	uscRepo := gormrepo.NewRepo(db)
	svcRepo := gormrepo.NewSvcRepo(db)
	usc := usecase.NewUsecase(uscRepo)
	w, p := uscqueue.NewQueueSys(uscqueue.Config{
		CreateAccQSize: 10000,
		FindAccQSize:   3000,
		ReadBatchSize:  100,
		WriteBatchSize: 2000,
		WaitTimeoutMs:  10,
	}, usc)
	svc := service.NewService(svcRepo)
	return &Container{
		QWorker:   w,
		QProducer: p,
		Service:   svc,
	}
}

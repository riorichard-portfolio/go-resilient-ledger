package account

import (
	"go-resilient-ledger/internal/account/gormrepo"
	"go-resilient-ledger/internal/account/uscqueue"
	"go-resilient-ledger/internal/account/usecase"

	"gorm.io/gorm"
)

type Container struct {
	gormRepo *gormrepo.GormRepo
	usc      *usecase.Usecase

	QWorker   *uscqueue.Worker
	QProducer *uscqueue.Producer
}

func NewContainer(
	db *gorm.DB,
) *Container {
	gormRepo := gormrepo.NewRepo(db)
	usc := usecase.NewUsecase(gormRepo)
	w, p := uscqueue.NewQueueSys(uscqueue.Config{
		CreateAccQSize: 10000,
		FindAccQSize:   3000,
		ReadBatchSize:  100,
		WriteBatchSize: 2000,
		WaitTimeoutMs:  10,
	}, usc)
	return &Container{
		gormRepo:  gormRepo,
		QWorker:   w,
		QProducer: p,
	}
}

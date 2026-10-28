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

func (c *Container) Usecase() *usecase.Usecase {
	if c.usc != nil {
		return c.usc
	}
	c.usc = usecase.NewUsecase(
		c.gormRepo,
	)
	return c.usc
}

func NewContainer(
	db *gorm.DB,
) *Container {
	gormRepo := gormrepo.NewRepo(db)
	w, p := uscqueue.NewQueueSys(uscqueue.Config{
		CreateAccQSize: 3000,
		FindAccQSize:   3000,
		ReadBatchSize:  100,
		WriteBatchSize: 500,
		WaitTimeoutMs:  10,
	})
	return &Container{
		gormRepo:  gormRepo,
		QWorker:   w,
		QProducer: p,
	}
}

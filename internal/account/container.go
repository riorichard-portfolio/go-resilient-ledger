package account

import (
	"go-resilient-ledger/internal/account/gormrepo"
	"go-resilient-ledger/internal/account/usecase"

	"gorm.io/gorm"
)

type Container struct {
	GormRepo *gormrepo.GormRepo

	usc *usecase.Usecase
}

func (c *Container) Usecase() *usecase.Usecase {
	if c.usc != nil {
		return c.usc
	}
	c.usc = usecase.NewUsecase(
		c.GormRepo,
	)
	return c.usc
}

func NewContainer(
	db *gorm.DB,
) *Container {
	gormRepo := gormrepo.NewRepo(db)
	return &Container{
		GormRepo: gormRepo,
	}
}

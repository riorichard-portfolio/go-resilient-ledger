package gormtx

import (
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"go-resilient-ledger/internal/account/shared"
	"go-resilient-ledger/internal/payment/gormrepo"
	"go-resilient-ledger/internal/payment/usecase"
	"go-resilient-ledger/internal/infra/gormdb"
)

type GormTxManager struct {
	db *gorm.DB
}

func NewGormTxManager(
	db *gorm.DB,
) *GormTxManager {
	return &GormTxManager{
		db: db,
	}
}

func (m *GormTxManager) New() (*usecase.Transaction, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "failed to begin gorm transaction")
	}
	return &usecase.Transaction{
		Repo:        gormrepo.NewRepo(tx),
		AccSvc:      shared.NewTxService(tx),
		Transaction: gormdb.NewTxBasic(tx),
	}, nil
}

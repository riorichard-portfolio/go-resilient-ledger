package shared

import (
	"gorm.io/gorm"

	"go-resilient-ledger/internal/account/gormrepo"
	"go-resilient-ledger/internal/account/service"
)

func NewTxService(
	tx *gorm.DB,
) *TxService {
	repo := gormrepo.NewSvcRepo(tx)
	return &TxService{
		service.NewService(repo),
	}
}

package gormrepo

import (
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"go-resilient-ledger/internal/account/service"
)

type SvcGormRepo struct {
	db *gorm.DB
}

func NewSvcRepo(
	db *gorm.DB,
) *SvcGormRepo {
	return &SvcGormRepo{
		db: db,
	}
}

func (gr *SvcGormRepo) BulkFindByIds(ids []string) (map[string]service.AccountMapVal, error) {
	accountsFound := make([]AccountGorm, 0, len(ids))
	err := gr.db.Where("id IN ?", ids).Find(&accountsFound).Error
	if err != nil {
		return map[string]service.AccountMapVal{}, errors.Wrap(err, "failed to bulk find accounts")
	}
	if len(accountsFound) < 1 {
		return map[string]service.AccountMapVal{}, nil
	}
	result := make(map[string]service.AccountMapVal, len(ids))
	for _, acc := range accountsFound {
		result[acc.ID] = service.AccountMapVal{
			Name:    acc.Name,
			Balance: acc.Balance,
		}
	}
	return result, nil
}

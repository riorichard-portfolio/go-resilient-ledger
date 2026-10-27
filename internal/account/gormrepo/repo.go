package gormrepo

import (
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"go-resilient-ledger/internal/account/entity"
	"go-resilient-ledger/internal/account/usecase"
)

type GormRepo struct {
	db *gorm.DB
}

func NewRepo(
	db *gorm.DB,
) *GormRepo {
	return &GormRepo{
		db: db,
	}
}

func (gr *GormRepo) BulkSaveNew(data []entity.Account) error {
	bulkData := make([]AccountGorm, 0, len(data))
	for _, acc := range data {
		bulkData = append(bulkData, AccountGorm{
			ID:      acc.ID(),
			Name:    acc.Name(),
			Balance: acc.Balance(),
		})
	}
	err := gr.db.Create(bulkData).Error
	if err != nil {
		return errors.Wrap(err, "failed to bulk save account")
	}
	return nil
}

func (gr *GormRepo) BulkFindByIds(ids []string) (map[string]usecase.AccountMapVal, error) {
	accountsFound := make([]AccountGorm, 0, len(ids))
	err := gr.db.Where("id IN ?", ids).Find(&accountsFound).Error
	if err != nil {
		return map[string]usecase.AccountMapVal{}, errors.Wrap(err, "failed to bulk find accounts")
	}
	if len(accountsFound) < 1 {
		return map[string]usecase.AccountMapVal{}, nil
	}
	result := make(map[string]usecase.AccountMapVal, len(ids))
	for _, acc := range accountsFound {
		result[acc.ID] = usecase.AccountMapVal{
			Name:    acc.Name,
			Balance: acc.Balance,
		}
	}
	return result, nil
}

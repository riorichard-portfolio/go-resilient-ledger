package usecase

import (
	"github.com/google/uuid"

	"go-resilient-ledger/internal/account/entity"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
)

type Repo interface {
	BulkSaveNew(data []entity.Account) error
	FindByIds(ids []string) []entity.Account
}

type Usecase struct {
	repo Repo
}

func NewUsecase(
	repo Repo,
) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) CreateAccount(data []createacc.WQdata) {
	bulkAccounts := make([]entity.Account, 0, len(data))
	successIndices := make([]int, 0, len(data))
	for i, item := range data {
		account, err := entity.NewAccount(
			uuid.NewString(),
			item.Name,
			item.Balance,
		)
		if err != nil {
			item.Reply <- err
		} else {
			bulkAccounts = append(bulkAccounts, account)
			successIndices = append(successIndices, i)
		}
	}
	if len(bulkAccounts) < 1 {
		return
	}
	err := u.repo.BulkSaveNew(bulkAccounts)
	for _, index := range successIndices {
		data[index].Reply <- err
	}
}

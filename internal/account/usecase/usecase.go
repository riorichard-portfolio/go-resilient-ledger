package usecase

import (
	"context"

	"github.com/google/uuid"

	"go-resilient-ledger/internal/account/entity"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
)

type Repo interface {
	BulkSaveNew(data []entity.Account) error
	BulkFindByIds(ids []string) (map[string]AccountMapVal, error)
	SaveNew(ctx context.Context, data entity.Account) error
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

func (u *Usecase) BulkCreateAccount(data []createacc.WQdata) {
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

func (u *Usecase) BulkFindAccounts(data []findacc.WQdata) {
	ids := make([]string, 0, len(data))
	seenIds := make(map[string]struct{}, len(data))
	for _, item := range data {
		_, exists := seenIds[item.ID]
		if !exists {
			ids = append(ids, item.ID)
			seenIds[item.ID] = struct{}{}
		}
	}
	accountMap, err := u.repo.BulkFindByIds(ids)
	for _, item := range data {
		accVal, exists := accountMap[item.ID]
		output := findacc.Output{
			Account: findacc.Account{
				ID:      item.ID,
				Name:    accVal.Name,
				Balance: accVal.Balance,
			},
			Err: nil,
		}
		if err != nil {
			output.Err = err
		} else if !exists {
			output.Err = AccountNotFoundErr
		}
		item.Reply <- output
	}
}

func (u *Usecase) IndividualCreateAccount(ctx context.Context, data createacc.Input) error {
	account, err := entity.NewAccount(
		uuid.NewString(),
		data.Name,
		data.Balance,
	)
	if err != nil {
		return err
	}
	err = u.repo.SaveNew(ctx, account)
	return err
}

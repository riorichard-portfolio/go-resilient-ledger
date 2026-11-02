package usecase

import (
	"github.com/google/uuid"

	"go-resilient-ledger/internal/payment/entity"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/createpaym"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/findpaym"
)

type Repo interface {
	BulkSaveNew(data []entity.Payment) error
	BulkFind(queries []BulkFindInput) map[string]PaymentMapVal
}

type AccountSvc interface {
	FindValidAccs(accIds []string) (map[string]struct{}, error)
}

type Usecase struct {
	repo   Repo
	accSvc AccountSvc
}

func NewUsecase(
	repo Repo,
) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) BulkCreatePayment(data []createpaym.WQdata) {
	payments := make([]entity.Payment, 0, len(data))
	accIds := make([]string, 0, len(data))
	successIndices := make([]int, 0, len(data))
	for i, d := range data {
		payment, err := entity.NewPayment(
			uuid.NewString(),
			d.AccountID,
			d.Amount,
		)
		if err != nil {
			d.Reply <- err
		} else {
			payments = append(payments, payment)
			accIds = append(accIds, d.AccountID)
			successIndices = append(successIndices, i)
		}
	}
	if len(payments) == 0 {
		return
	}
	validAccsMap, err := u.accSvc.FindValidAccs(accIds)
	if err != nil {
		for _, si := range successIndices {
			data[si].Reply <- err
		}
		return
	}
	validIdx := 0
	for i, p := range payments {
		_, isValidAcc := validAccsMap[p.AccountID()]
		if isValidAcc {
			payments[validIdx] = p
			successIndices[validIdx] = successIndices[i]
			validIdx++
			continue
		}
		data[successIndices[i]].Reply <- InvalidAccErr
	}
	if validIdx == 0 {
		return
	}
	payments = payments[:validIdx]
	successIndices = successIndices[:validIdx]
	err = u.repo.BulkSaveNew(payments)
	for _, si := range successIndices {
		data[si].Reply <- err
	}
}

func (u *Usecase) BulkFindPayments(data []findpaym.WQdata) {

}

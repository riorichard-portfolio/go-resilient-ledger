package usecase

import (
	"time"

	"github.com/google/uuid"

	"go-resilient-ledger/internal/payment/entity"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/createpaym"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/findpaym"
	txpkg "go-resilient-ledger/internal/pkg/transaction"
)

type Repo interface {
	BulkFind(queries []BulkFindInput) (map[int][]PaymentMapVal, error)
}

type AccountSvc interface {
	FindValidAccs(accIds []string) (map[string]struct{}, error)
}

type TxRepo interface {
	BulkInsert(data []entity.Payment) error
}

type TxAccountSvc interface {
	AddAccsBalance(bmap map[string]int) error
}

type Transaction struct {
	Repo   TxRepo
	AccSvc TxAccountSvc
	txpkg.Transaction
}

type TxManager interface {
	New() (*Transaction, error)
}

type Usecase struct {
	repo      Repo
	accSvc    AccountSvc
	txManager TxManager
}

func NewUsecase(
	repo Repo,
	accSvc AccountSvc,
	txManager TxManager,
) *Usecase {
	return &Usecase{
		repo:      repo,
		accSvc:    accSvc,
		txManager: txManager,
	}
}

func (u *Usecase) BulkCreatePayment(data []createpaym.WQdata) {
	payments := make([]entity.Payment, 0, len(data))
	accIds := make([]string, 0, len(data))
	successIndices := make([]int, 0, len(data))
	now := time.Now().UTC()
	for i, d := range data {
		payment, err := entity.NewPayment(
			uuid.NewString(),
			d.AccountID,
			d.Amount,
			now,
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
	balanceMap := make(map[string]int, len(payments))
	for _, p := range payments {
		balanceMap[p.AccountID()] += p.Amount()
	}
	tx, err := u.txManager.New()
	if err != nil {
		for _, si := range successIndices {
			data[si].Reply <- err
		}
		return
	}
	defer tx.Rollback()
	err = tx.Repo.BulkInsert(payments)
	if err != nil {
		for _, si := range successIndices {
			data[si].Reply <- err
		}
		return
	}
	err = tx.AccSvc.AddAccsBalance(balanceMap)
	if err != nil {
		for _, si := range successIndices {
			data[si].Reply <- err
		}
		return
	}
	tx.Commit()
	for _, si := range successIndices {
		data[si].Reply <- nil
	}
}

func (u *Usecase) BulkFindPayments(data []findpaym.WQdata) {
	queries := make([]BulkFindInput, 0, len(data))
	for i, d := range data {
		queries = append(queries, BulkFindInput{
			AccountID: d.AccountID,
			QueryID:   i,
			Filter:    d.Filter,
		})
	}
	paymentsReqMap, err := u.repo.BulkFind(queries)
	if err != nil {
		for _, d := range data {
			d.Reply <- findpaym.Output{
				Err: err,
			}
		}
		return
	}
	for i, d := range data {
		payments := paymentsReqMap[i]
		paymentsReply := make([]findpaym.Payment, 0, len(payments))
		for _, p := range payments {
			paymentsReply = append(paymentsReply, findpaym.Payment{
				ID:        p.ID,
				AccountID: p.AccountID,
				Amount:    p.Amount,
				CreatedAt: p.CreatedAt,
			})
		}
		d.Reply <- findpaym.Output{
			Err:      nil,
			Payments: paymentsReply,
		}
	}
}

package gormrepo

import (
	"fmt"
	"strings"

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

func (gr *SvcGormRepo) BulkAddBalance(bmap map[string]int) error {
	// Jika map kosong, langsung return tanpa eksekusi query
	if len(bmap) == 0 {
		return nil
	}

	var valueRows []string
	var args []interface{}

	// Iterasi map untuk menyiapkan baris nilai dan parameter binding
	for accID, delta := range bmap {
		valueRows = append(valueRows, "(?, ?::int)") // Cast tipe data integer jika diperlukan
		args = append(args, accID, delta)
	}

	// Susun query PostgreSQL menggunakan VALUES table expression
	// Asumsi nama tabel di database adalah "accounts", dengan kolom "id" dan "balance"
	query := fmt.Sprintf(`
		UPDATE accounts AS a 
		SET balance = a.balance + v.delta 
		FROM (VALUES %s) AS v(id, delta) 
		WHERE a.id = v.id
	`, strings.Join(valueRows, ", "))

	// Eksekusi secara aman dengan GORM db.Exec (1 roundtrip)
	return gr.db.Exec(query, args...).Error
}
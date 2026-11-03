package gormrepo

import (
	"strings"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"go-resilient-ledger/internal/payment/entity"
	"go-resilient-ledger/internal/payment/usecase"
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

func (gr *GormRepo) BulkSaveNew(data []entity.Payment) error {
	bulkData := make([]PaymentGorm, 0, len(data))
	for _, p := range data {
		bulkData = append(bulkData, PaymentGorm{
			ID:        p.ID(),
			AccountID: p.AccountID(),
			Amount:    p.Amount(),
			CreatedAt: p.CreatedAt(),
		})
	}
	err := gr.db.Create(bulkData).Error
	if err != nil {
		return errors.Wrap(err, "failed to bulk save payment")
	}
	return nil
}

func (gr *GormRepo) BulkFind(queries []usecase.BulkFindInput) (map[int][]usecase.PaymentMapVal, error) {
	if len(queries) == 0 {
		return make(map[int][]usecase.PaymentMapVal), nil
	}
	values := make([]string, 0, len(queries))
	args := make([]interface{}, 0, 5*len(queries))

	for _, q := range queries {
		values = append(values, "(?::int, ?::text, ?::timestamptz, ?::timestamptz, ?::int)")

		args = append(args,
			q.QueryID,
			q.AccountID,
			q.From,
			q.Until,
			q.Limit,
		)
	}

	query := `
        WITH requests(query_id, account_id, from_time, until_time, row_limit) AS (
            VALUES ` + strings.Join(values, ",") + `
        )
        SELECT
            r.query_id,
        	 p.id,
            p.account_id,
            p.amount,
            p.created_at
        FROM requests r
        LEFT JOIN LATERAL (
            SELECT id, account_id, amount, created_at
            FROM payments
            WHERE account_id = r.account_id
              AND created_at >= r.from_time
              AND created_at <= r.until_time
            ORDER BY created_at DESC
            LIMIT r.row_limit
        ) p ON TRUE
    `
	var dbResults []BulkFindRes

	err := gr.db.Raw(query, args...).Scan(&dbResults).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to bulk find payments")
	}
	result := make(map[int][]usecase.PaymentMapVal, len(queries))
	for _, dr := range dbResults {
		if dr.ID == "" {
			continue
		}
		result[dr.QueryID] = append(result[dr.QueryID], usecase.PaymentMapVal{
			ID:        dr.ID,
			AccountID: dr.AccountID,
			Amount:    dr.Amount,
			CreatedAt: dr.CreatedAt,
		})
	}
	return result, nil
}

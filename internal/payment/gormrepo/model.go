package gormrepo

import "time"

type PaymentGorm struct {
	ID        string    `gorm:"type:uuid;column:id;primaryKey"`
	AccountID string    `gorm:"type:varchar(255);column:account_id;not null"`
	Amount    int       `gorm:"type:int;column:amount;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
}

type BulkFindRes struct {
	QueryID int `gorm:"type:int;column:query_id;not null"`
	PaymentGorm
}

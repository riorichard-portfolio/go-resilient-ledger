package gormdb

import (
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type TxBasic struct {
	tx *gorm.DB
}

func NewTxBasic(
	tx *gorm.DB,
) *TxBasic {
	return &TxBasic{
		tx: tx,
	}
}

func (t *TxBasic) Commit() error {
	err := t.tx.Commit().Error
	if err != nil {
		return errors.Wrap(err, "failed to commit gorm transaction")
	}
	return nil
}

func (t *TxBasic) Rollback() error {
	err := t.tx.Rollback().Error
	if err != nil {
		return errors.Wrap(err, "failed to rollback gorm transaction")
	}
	return nil
}

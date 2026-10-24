package gormdb

import (
	"time"

	"go-resilient-ledger/internal/infra/config"

	"github.com/cockroachdb/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(
	cfg config.PostgreConfig,
) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to open gorm connection")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get underlying sql.DB")
	}

	sqlDB.SetMaxIdleConns(10)                  // Jumlah koneksi idle di pool
	sqlDB.SetMaxOpenConns(100)                 // Batas maksimal koneksi simultan ke Postgres
	sqlDB.SetConnMaxLifetime(1 * time.Hour)    // Umur maksimal satu koneksi sebelum didaur ulang
	sqlDB.SetConnMaxIdleTime(15 * time.Minute) // Durasi koneksi idle ditutup dari pool

	err = sqlDB.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping DB")
	}

	return db, nil
}

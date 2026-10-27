package usecase

import "go-resilient-ledger/internal/pkg/apperr"

var(
	AccountNotFoundErr = apperr.NotFoundErr("account not found")
)
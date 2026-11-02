package usecase

import "go-resilient-ledger/internal/pkg/apperr"

var(
	InvalidAccErr = apperr.UnauthorizedErr("invalid account")
)
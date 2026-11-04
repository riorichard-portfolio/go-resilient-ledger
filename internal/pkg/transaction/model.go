package transaction

import "context"

type Transaction interface {
	Begin() error
	Commit() error
	Rollback() error
}

type CtxTransaction interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
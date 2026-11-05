package transaction

import "context"

type Transaction interface {
	Commit() error
	Rollback() error
}

type CtxTransaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
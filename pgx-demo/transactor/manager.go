package transactor

import (
	"context"
	"database/sql"
	"errors"
)

// Manager wraps default sql.Tx
type Manager interface {
	// Process starts a new transaction, processes txFunc and
	// if txFunc returns error, then tries to rollback the transaction
	// if txFunc returns nil, then tries to commit the transaction
	// if the func couldn't commit or rollback the transaction, will return error
	// if panic is occurred in txFunc, the transaction will be rollbacked
	Process(ctx context.Context, txFunc func(ctx context.Context) error, opts ...sql.TxOptions) error
}

// ErrTxCommitRollback occurs when an error has occurred in a transaction and
// Commit() is called. DB accepts COMMIT on aborted transactions, but
// it is treated as ROLLBACK.
var ErrTxCommitRollback = errors.New("commit unexpectedly resulted in rollback")

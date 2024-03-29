package transactor

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/joomcode/errorx"
	"go.uber.org/zap"
)

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

type TxManager struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewTxManager(db *sql.DB, logger *zap.Logger) *TxManager {
	return &TxManager{db: db, logger: logger}
}

func (m TxManager) Process(ctx context.Context, txFunc func(tx context.Context) error, opts ...sql.TxOptions) (errTx error) {
	// check that tx has been injected before
	if tx := extractTx(ctx); tx != nil {
		return txFunc(ctx)
	}

	var opt *sql.TxOptions
	if len(opts) > 0 {
		opt = &opts[0]
	}

	sqlTx, err := m.db.BeginTx(ctx, opt)
	if err != nil {
		return errorx.InternalError.Wrap(err, "couldn't init sql transaction")
	}

	defer func() {
		if r := recover(); r != nil {
			errTx = errorx.InternalError.New("recovered panic in a transaction: %+v", r)
			m.logger.Error("recovered error in a transaction", zap.Any("panic", r))
			if err := sqlTx.Rollback(); err != nil {
				m.logger.Error("couldn't rollback transaction in panic", zap.Error(err))
			}
		}
	}()

	// run callback
	err = txFunc(injectTx(ctx, sqlTx))
	if err != nil {
		// if error, rollback
		if errRollback := sqlTx.Rollback(); errRollback != nil {
			if errors.Is(errRollback, sql.ErrTxDone) {
				m.logger.Info("couldn't rollback transaction", zap.Error(err), zap.Error(errRollback))
				return err
			}
			m.logger.Error("couldn't rollback transaction", zap.Error(err), zap.Error(errRollback))
			return errorx.InternalError.Wrap(errRollback, "couldn't rollback transaction")
		}
		return err
	}

	// if no error, commit
	if errCommit := sqlTx.Commit(); errCommit != nil {
		if errors.Is(err, pgx.ErrTxCommitRollback) {
			return ErrTxCommitRollback
		}
		m.logger.Error("couldn't commit transaction", zap.Error(errCommit))
		return errorx.InternalError.Wrap(errCommit, "couldn't commit transaction")
	}
	return nil
}

package transactor

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type DBTXWrapper struct {
	db *sql.DB
}

func NewTxWrapper(db *sql.DB) *DBTXWrapper {
	return &DBTXWrapper{db: db}
}

func (w *DBTXWrapper) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx := extractTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return w.db.ExecContext(ctx, query, args...)
}

func (w *DBTXWrapper) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if tx := extractTx(ctx); tx != nil {
		return tx.PrepareContext(ctx, query)
	}
	return w.db.PrepareContext(ctx, query)
}

func (w *DBTXWrapper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if tx := extractTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}
	return w.db.QueryContext(ctx, query, args...)
}

func (w *DBTXWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if tx := extractTx(ctx); tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return w.db.QueryRowContext(ctx, query, args...)
}

type txKey struct{}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}

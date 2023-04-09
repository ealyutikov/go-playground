package repo

import (
	"context"

	"github.com/elyutikov/pgx-demo/domain"
	"github.com/elyutikov/pgx-demo/transactor"
)

type OutboxRepo struct {
	tx transactor.DBTX
}

func NewOutboxRepo(tx transactor.DBTX) *OutboxRepo {
	return &OutboxRepo{tx: tx}
}

const saveOutboxSQL = "insert into outbox(data, created_at) values ($1,$2);"

func (repo *OutboxRepo) Save(ctx context.Context, o domain.Outbox) error {
	if _, err := repo.tx.ExecContext(ctx, saveOutboxSQL, o.Data, o.CreatedAt); err != nil {
		return err
	}

	return nil
}

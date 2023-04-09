package repo

import (
	"context"

	"github.com/elyutikov/pgx-demo/domain"
	"github.com/elyutikov/pgx-demo/transactor"
)

type UserRepo struct {
	tx transactor.DBTX
}

func NewUserRepo(tx transactor.DBTX) *UserRepo {
	return &UserRepo{tx: tx}
}

const saveUserSQL = "insert into users values($1,$2,$3,$4);"

func (u *UserRepo) Save(ctx context.Context, user domain.User) error {
	if _, err := u.tx.ExecContext(ctx, saveUserSQL, user.ID, user.Name, user.IsActive, user.CreatedAt); err != nil {
		return err
	}
	return nil
}

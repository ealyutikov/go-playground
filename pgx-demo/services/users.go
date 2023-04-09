package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elyutikov/pgx-demo/domain"
	"github.com/elyutikov/pgx-demo/repo"
	"github.com/elyutikov/pgx-demo/transactor"
)

type UserService struct {
	txManager transactor.Manager
	uRepo     *repo.UserRepo
	oRepo     *repo.OutboxRepo
}

func NewUserService(txManager transactor.Manager, uRepo *repo.UserRepo, oRepo *repo.OutboxRepo) *UserService {
	return &UserService{txManager: txManager, uRepo: uRepo, oRepo: oRepo}
}

func (u *UserService) Save(ctx context.Context, user domain.User) error {
	userJson, _ := json.Marshal(user)

	if err := u.txManager.Process(ctx, func(ctx context.Context) error {
		// save to users
		if err := u.uRepo.Save(ctx, user); err != nil {
			return err
		}

		// save to outbox
		if err := u.oRepo.Save(ctx, domain.Outbox{Data: userJson, CreatedAt: time.Now()}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return fmt.Errorf("got error in the transaction during saving user: %w", err)
	}

	return nil
}

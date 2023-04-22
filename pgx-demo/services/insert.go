package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/elyutikov/pgx-demo/domain"
	"github.com/elyutikov/pgx-demo/repo"
	"github.com/elyutikov/pgx-demo/transactor"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func InsertDemo(sqldb *sql.DB) error {
	dbtx := transactor.NewTxWrapper(sqldb)

	uRepo := repo.NewUserRepo(dbtx)
	oRepo := repo.NewOutboxRepo(dbtx)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	tx := transactor.NewTxManager(sqldb, logger)
	uService := NewUserService(tx, uRepo, oRepo)

	user := domain.User{
		ID:        uuid.New(),
		Name:      "Alex",
		IsActive:  true,
		CreatedAt: time.Now(),
	}

	return uService.Save(context.TODO(), user)
}

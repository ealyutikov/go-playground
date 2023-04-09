package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/elyutikov/pgx-demo/db"
	"github.com/elyutikov/pgx-demo/domain"
	"github.com/elyutikov/pgx-demo/repo"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func InsertDemo(sqldb *sql.DB) error {
	dbtx := db.NewTxWrapper(sqldb)

	uRepo := repo.NewUserRepo(dbtx)
	oRepo := repo.NewOutboxRepo(dbtx)

	tx := db.NewTxManager(sqldb, zap.NewNop())
	uService := NewUserService(tx, uRepo, oRepo)

	user := domain.User{
		ID:        uuid.New(),
		Name:      "Alex",
		IsActive:  true,
		CreatedAt: time.Now(),
	}

	return uService.Save(context.TODO(), user)
}

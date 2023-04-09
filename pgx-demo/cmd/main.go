package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/elyutikov/pgx-demo/services"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func main() {
	connectionString := "postgres://postgres:postgres@localhost:5432/demo"
	migrationsPath := "./migrations"

	sqldb, err := newSQLDB(connectionString)
	if err != nil {
		panic(err)
	}

	if err = runMigrations(sqldb, migrationsPath); err != nil {
		panic(err)
	}

	if err = services.InsertDemo(sqldb); err != nil {
		panic(err)
	}
}

func newSQLDB(connectionString string) (*sql.DB, error) {
	config, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse postgres connection string: %w", err)
	}

	db := stdlib.OpenDB(*config)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetConnMaxIdleTime(3 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to connect to postgres: %w", err)
	}

	return db, nil
}

func runMigrations(sqldb *sql.DB, path string) error {
	if err := goose.Up(sqldb, path); err != nil {
		return fmt.Errorf("unable to execute migration: %w", err)
	}
	return nil
}

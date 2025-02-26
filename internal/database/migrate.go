package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/iremsha/oapicodegen-example/internal/config"
	logger "github.com/iremsha/oapicodegen-example/internal/log"
)

func MigrateUp(cfg config.Database, log *logger.Logger, migrationsPath string) error {
	ctx := context.Background()
	dsn := cfg.GetMigrateDsn()
	if dsn == "" {
		return errors.New("migrate: environment variable not declared: PG_URL")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	schemaSQL := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS \"%s\"", cfg.Schema)
	if _, err := db.Exec(schemaSQL); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	attempts := 5
	var m *migrate.Migrate
	for attempts > 0 {
		m, err = migrate.New(fmt.Sprintf("file://%s", migrationsPath), dsn)
		if err == nil {
			break
		}

		log.Info(ctx, fmt.Sprintf("Migrate: postgres is trying to connect %s , attempts left: %d", dsn, attempts))
		time.Sleep(1 * time.Second)
		attempts--
	}

	if err != nil {
		return fmt.Errorf("failed to initiate migration: %w", err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Error(ctx, err.Error())
		return fmt.Errorf("migration up error: %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Info(ctx, "Migrate: no change")
	} else {
		log.Info(ctx, "Migrate: up success")
	}

	return nil
}

func Fixtures(cfg config.Database, fixtureSQL string) error {
	db, err := sql.Open("postgres", cfg.GetMigrateDsn())
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	if _, err := db.Exec(fixtureSQL); err != nil {
		return fmt.Errorf("failed to execute fixtures: %w", err)
	}

	return nil
}

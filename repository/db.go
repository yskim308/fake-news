package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func (r *Repository) Connect(
	ctx context.Context,
) error {
	dbConfig := Config{
		Host:     getEnvOrThrow("CLUSTER_ENDPOINT"),
		Region:   getEnvOrThrow("REGION"),
		Port:     getEnv("DB_PORT", "5432"),
		Database: getEnv("DB_NAME", "postgres"),
		Password: "",
	}

	url := CreateConnectionString(dbConfig)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return fmt.Errorf("unable to parse pool config")
	}

	poolConfig.BeforeConnect = func(ctx context.Context, cfg *pgx.ConnConfig) error {
		token, err := GenerateDbConnectToken(ctx, dbConfig.Host, dbConfig.Region)
		if err != nil {
			return fmt.Errorf("failed to generate auth token %w", err)
		}
		cfg.Password = token
		return nil
	}

	poolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		schema := "public"
		_, err := conn.Exec(ctx, fmt.Sprintf("SET search_path = %s", schema))
		if err != nil {
			return fmt.Errorf("failed to set search_path to %s: %w", schema, err)
		}
		return nil
	}

	pgxPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %v", err)
	}

	r.db = pgxPool
	return nil
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) Connect() {
	ctx := context.Background()

	endpoint := os.Getenv("DSQL_ENDPOINT")
	region := os.Getenv("AWS_REGION")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	if endpoint == "" || region == "" || dbUser == "" || dbName == "" {
		log.Fatal("Missing env vars: DSQL_ENDPOINT, AWS_REGION, DB_USER, DB_NAME")
	}

	// Load AWS credentials (Lambda automatically injects them)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	// Generate IAM auth token (the password)
	rdsClient := rds.NewFromConfig(cfg)
	authToken, err := rds.BuildAuthToken(
		ctx,
		endpoint,
		region,
		dbUser,
		cfg.Credentials,
	)
	if err != nil {
		log.Fatalf("failed to generate auth token: %v", err)
	}

	// Build connection string
	dsn := fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=require",
		endpoint, dbUser, authToken, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("open DB error: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("ping DB error: %v", err)
	}

	r.db = db
	fmt.Println("✅ Connected to Aurora DSQL")
}

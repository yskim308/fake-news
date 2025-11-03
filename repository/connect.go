package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dsql/auth"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Region   string
}

func GenerateDbConnectToken(
	ctx context.Context, clusterEndpoint, region string,
) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	//assuming that the user is admin
	token, err := auth.GenerateDBConnectAdminAuthToken(ctx, clusterEndpoint, region, cfg.Credentials)
	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateConnectionString(dbConfig Config) string {
	var str strings.Builder
	str.WriteString("postgres://")
	str.WriteString(dbConfig.User)
	str.WriteString("@")
	str.WriteString(dbConfig.Host)
	str.WriteString(":")
	str.WriteString(dbConfig.Port)
	str.WriteString("/")
	str.WriteString(dbConfig.Database)
	str.WriteString("?sslmode=verify-full")
	str.WriteString("&sslnegotiation=direct")
	url := str.String()
	return url
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrThrow(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Errorf("environment variable %s not set", key))
	}
	return value
}

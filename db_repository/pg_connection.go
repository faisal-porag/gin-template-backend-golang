package db_repository

import (
	"context"
	"fmt"
	"github.com/faisal-porag/gin-template-backend-golang/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type PgRepository struct {
	db *gorm.DB
}

// NewPgRepository ->>>> POSTGRES DATABASE CONNECT SETUP
func NewPgRepository(cfg *config.Config) (*PgRepository, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		cfg.PGDatabaseHost, cfg.PGDatabasePort, cfg.PGDatabaseUsername, cfg.PGDatabasePassword, cfg.PGDatabaseName,
		cfg.PGTimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to establish a connection to postgres")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("failed to get pg db instance")
		return nil, err
	}

	// IF GETTING 0
	if cfg.PGDatabaseConnectionTimeout == 0 {
		cfg.PGDatabaseConnectionTimeout = time.Second * 180 // 3 min
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), cfg.PGDatabaseConnectionTimeout)
	defer pingCancel()

	// Ping to check if the database connection is active
	if err := sqlDB.PingContext(pingCtx); err != nil {
		log.Error().Err(err).Msg("Postgres SQL is down! please check")
		return nil, err
	}

	return &PgRepository{
		db: db,
	}, nil
}

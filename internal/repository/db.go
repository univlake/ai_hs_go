package repository

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ai_hs_go/internal/config"
)

func Open(cfg *config.Config, log *zap.Logger) (*gorm.DB, error) {
	gormLogLevel := logger.Warn
	if cfg.App.Env == "local" {
		gormLogLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.New(zapWriter{log: log}, logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  gormLogLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		}),
		// PrepareStmt is intentionally off: gorm-adapter v3 reuses the same
		// connection pool and its session-level mutations (e.g. table creates)
		// can clash with cached prepared statements bound to the prior catalog
		// state on PostgreSQL.
	})
	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("gorm sqlDB: %w", err)
	}
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime())

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return db, nil
}

type zapWriter struct {
	log *zap.Logger
}

func (w zapWriter) Printf(format string, args ...interface{}) {
	w.log.Sugar().Infof(format, args...)
}

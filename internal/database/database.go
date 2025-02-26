package database

import (
	"time"

	"github.com/iremsha/oapicodegen-example/internal/config"
	logger "github.com/iremsha/oapicodegen-example/internal/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func New(cfg config.Config, log *logger.Logger) (*Database, error) {
	newLogger := glog.New(
		&logger.SlogAdapter{Logger: log},
		glog.Config{
			SlowThreshold: 10 * time.Second,
			LogLevel:      glog.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDsn()), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
		TranslateError:         true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)

	return &Database{
		DB: db,
	}, nil
}

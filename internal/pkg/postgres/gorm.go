package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dinosgnk/agora-project/internal/pkg/config"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
)

type GormDatabase struct {
	*gorm.DB
}

func NewGormDatabase(log logger.Logger, gormCfg *gorm.Config) (*GormDatabase, error) {
	postgresCfg := config.LoadConfig[PostgresConfig](log)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		postgresCfg.User,
		postgresCfg.Password,
		postgresCfg.Host,
		postgresCfg.Port,
		postgresCfg.DBName,
	)

	log.Info("Connecting to PostgreSQL database",
		"host", postgresCfg.Host,
		"port", postgresCfg.Port,
		"database", postgresCfg.DBName,
		"user", postgresCfg.User)

	gormDb, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		log.Error("Failed to connect to PostgreSQL database", "error", err.Error())
		return nil, err
	}

	log.Info("Successfully connected to PostgreSQL database")

	db := &GormDatabase{DB: gormDb}
	return db, nil
}

func (gormDb *GormDatabase) GetDB() *gorm.DB {
	return gormDb.DB
}

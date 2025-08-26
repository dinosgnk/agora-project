package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dinosgnk/agora-project/pkg/config"
)

type GormDatabase struct {
	*gorm.DB
}

func NewGormDatabase(gormCfg *gorm.Config) (*GormDatabase, error) {
	postgresCfg := config.LoadConfig[PostgresConfig]()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		postgresCfg.User,
		postgresCfg.Password,
		postgresCfg.Host,
		postgresCfg.Port,
		postgresCfg.DBName,
	)

	gormDb, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		return nil, err
	}

	db := &GormDatabase{DB: gormDb}

	return db, nil
}

func (gormDb *GormDatabase) GetDB() *gorm.DB {
	return gormDb.DB
}

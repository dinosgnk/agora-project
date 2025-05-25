package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type DatabaseConfig struct {
// 	Host     string `mapstructure:"host"`
// 	Port     int    `mapstructure:"port"`
// 	User     string `mapstructure:"user"`
// 	Password string `mapstructure:"password"`
// 	DBName   string `mapstructure:"dbName"`
// 	SSLMode  bool   `mapstructure:"sslMode"`
// }

type IDatabase interface {
	GetDB() *gorm.DB
}

type Database struct {
	gormDb *gorm.DB
	// config *DatabaseConfig
}

// func NewDatabase(cfg *DatabaseConfig) (*Database, error) {
func NewDatabase() (*Database, error) {
	// datasource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
	// 	cfg.User,
	// 	cfg.Password,
	// 	cfg.Host,
	// 	cfg.Port,
	// 	cfg.DBName,
	// )

	datasource := "postgres://admin:admin@localhost:3301/catalpg?sslmode=disable"

	gormDb, err := gorm.Open(postgres.Open(datasource), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Set up connection pool
	sqlDB, err := gormDb.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)

	return &Database{
		gormDb: gormDb,
	}, nil
}

func (db *Database) GetDB() *gorm.DB {
	return db.gormDb
}

// func (db *Database) Close() {
// 	sqlDB := db.GetDB()
// 	_ = sqlDB.Close()
// }

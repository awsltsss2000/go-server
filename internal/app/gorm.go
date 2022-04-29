package app

import (
	"go-server/internal/app/admin/dao/user"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := NewGormDB()
	if err != nil {
		panic(err)
	}

	// if err := autoMigrate(db); err != nil {
	// 	panic(err)
	// }
	return db
}

func NewGormDB() (*gorm.DB, error) {
	cfg := GetConfig()

	// gConfig := &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   cfg.Gorm.TablePrefix,
	// 		SingularTable: true,
	// 	},
	// }
	db, err := gorm.Open(postgres.Open(cfg.Postgres.Dsn))
	if err != nil {
		panic(err)
	}
	if cfg.RunMode == "debug" {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(cfg.Gorm.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Gorm.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Gorm.MaxLifetime) * time.Second)

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		new(user.User),
	)
}

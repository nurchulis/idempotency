package bootstrap

import (
	"fmt"
	"github.com/nurchulis/idempotency/config"
	"github.com/nurchulis/idempotency/database/seeder"
	"github.com/nurchulis/idempotency/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RegistrySQL(conf config.DBConfig) (*gorm.DB, error) {
	var (
		dbURL string
	)

	switch conf.Driver {
	case "mysql":
		dbURL = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName,
		)

		db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	return nil, fmt.Errorf("driver not found")
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
	)
}

func AutoSeed(db *gorm.DB) error {
	if err := seeder.UserSeeder(db); err != nil {
		return err
	}
	return nil
}

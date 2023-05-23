package seeder

import (
	"github.com/nurchulis/idempotency/internal/entity"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) error {
	// force delete all data
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		return err
	}
	// reset auto increment
	if err := db.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error; err != nil {
		return err
	}

	users := []entity.User{
		{
			Name:  "John Doe",
			Email: "johndoe@mail.com",
		},
		{
			Name:  "",
			Email: "nurchulis@mail.com",
		},
	}

	if err := db.Create(&users).Error; err != nil {
		return err
	}

	return nil
}

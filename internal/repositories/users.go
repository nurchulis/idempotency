package repositories

import (
	"github.com/nurchulis/idempotency/internal/entity"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (u *userRepo) FindAll() ([]entity.User, error) {
	var users []entity.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

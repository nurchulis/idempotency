package repositories

import "github.com/nurchulis/idempotency/internal/entity"

type UserRepo interface {
	FindAll() ([]entity.User, error)
}

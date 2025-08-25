package service

import (
	"github.com/ZephLevy/Safe-return-backend/internal/db"
)

const (
	HashCost = 12
)

type UserService struct {
	repo *db.UserRepository
}

func NewUserService(ur *db.UserRepository) *UserService {
	return &UserService{repo: ur}
}

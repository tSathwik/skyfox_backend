package user

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db:db}
}

func (r *userRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}
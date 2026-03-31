package user

import (
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"go.uber.org/zap"

)

type UserRepository interface {
	CreateUser(user *User) error
	GetUserById(id string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
	logger *zap.Logger
	
}

func NewUserRepository(db *gorm.DB, logger *zap.Logger) UserRepository {
	return &userRepository{db: db, logger: logger}
}

func (r *userRepository) CreateUser(user *User) error {
	id, err := uuid.NewV4()
	if err != nil {
		r.logger.Error("Failed to generate user ID", zap.Error(err))
		return err
	}
	user.ID = id.String()
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		r.logger.Error("Failed to hash password", zap.Error(err))
		return err
	}
	user.Password = string(hashedPassword)
	r.logger.Info("User created successfully", zap.String("user_id", user.ID))
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserById(id string) (*User, error) {
	var user User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		r.logger.Error("Failed to get user by id", zap.String("user_id", id), zap.Error(err))
		return nil, err
	}
	r.logger.Info("User retrieved successfully", zap.String("user_id", user.ID))
	return &user, nil
}
package user

import (
	"errors"
	"skyfox_backend/internal/user/dto"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(user *dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUserById(id string) (*dto.UserResponse, error)
}

type userService struct {
	repo UserRepository
	logger *zap.Logger
}

var (
	ErrInvalidUserId = errors.New("invalid user id")
)

func NewUserService(repo UserRepository, logger *zap.Logger) UserService {
	return &userService{repo: repo, logger: logger}
}

func (u *userService) CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error) {

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}

	err := u.repo.CreateUser(user)
	if err != nil {
		u.logger.Error("Failed to create user", zap.Error(err))
		return nil, err
	}
	u.logger.Info("User created successfully", zap.String("user_id", user.ID))
	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}, nil
}

func (u *userService) GetUserById(id string) (*dto.UserResponse, error) {
	if id == "" {
		u.logger.Error("Invalid user id provided")
		return nil, ErrInvalidUserId
	}

	user, err := u.repo.GetUserById(id)
	if err != nil {
		u.logger.Error("Failed to get user by id", zap.String("user_id", id), zap.Error(err))
		return nil, err
	}
	u.logger.Info("User retrieved successfully", zap.String("user_id", user.ID))

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}, nil
}	
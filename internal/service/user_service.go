package service

import (
	"context"

	"echo-mysql/internal/model"

	"gorm.io/gorm"
)

type UserService interface {
	All(ctx context.Context) (*[]model.User, error)
	Load(ctx context.Context, id string) (*model.User, error)
	Insert(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}

type userService struct {
	DB *gorm.DB
}

func (s *userService) All(ctx context.Context) (*[]model.User, error) {
	var users *[]model.User
	_ = s.DB.Find(&users)
	return users, nil
}

func (s *userService) Load(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	s.DB.First(&user, "id = ?", id)
	return &user, nil
}

func (s *userService) Insert(ctx context.Context, user *model.User) (*model.User, error) {
	s.DB.Create(&user)
	return user, nil
}

func (s *userService) Update(ctx context.Context, user *model.User) error {
	if err := s.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	if err := s.DB.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

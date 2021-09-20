package service

import (
	"context"
	"github.com/ZhansultanS/myLMS/backend/internal/dto"
	"github.com/ZhansultanS/myLMS/backend/internal/model"
	"github.com/ZhansultanS/myLMS/backend/internal/repository"
	"github.com/ZhansultanS/myLMS/backend/pkg/hasher"
	"strings"
)

type UserService struct {
	repo        repository.IUserRepository
	roleService IRoleService
	hasher      hasher.PasswordHasher
}

func (u *UserService) Create(ctx context.Context, userDto dto.UserCreateDto) error {
	hashedPassword, err := u.hasher.HashPassword(userDto.Password)
	if err != nil {
		return err
	}

	role, err := u.roleService.FindByName(ctx, userDto.RoleName)
	if err != nil {
		return err
	}

	profile := &model.Profile{
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
	}
	user := &model.User{
		Username:     strings.ToLower(userDto.Firstname[:1]) + strings.ToLower(userDto.Lastname),
		Profile:      profile,
		PasswordHash: hashedPassword,
		Role:         role,
	}

	return u.repo.Insert(ctx, user)
}

func (u *UserService) FindAll(ctx context.Context) ([]*model.User, error) {
	return u.repo.GetAll(ctx)
}

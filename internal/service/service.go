package service

import (
	"context"
	"github.com/ZhansultanS/myLMS/backend/internal/dto"
	"github.com/ZhansultanS/myLMS/backend/internal/model"
	"github.com/ZhansultanS/myLMS/backend/internal/repository"
	"github.com/ZhansultanS/myLMS/backend/pkg/hasher"
)

type Service struct {
	Roles IRoleService
	Users IUserService
}

type Deps struct {
	Repositories   repository.Repository
	PasswordHasher hasher.PasswordHasher
}

func NewService(deps Deps) Service {
	roles := &RoleService{repo: deps.Repositories.Roles}
	users := &UserService{repo: deps.Repositories.Users, roleService: roles, hasher: deps.PasswordHasher}
	return Service{
		Roles: roles,
		Users: users,
	}
}

type IRoleService interface {
	FindByName(ctx context.Context, name string) (*model.Role, error)
	FindAll(ctx context.Context) ([]*model.Role, error)
}

type IUserService interface {
	Create(ctx context.Context, userDto dto.UserCreateDto) error
	FindAll(ctx context.Context) ([]*model.User, error)
}

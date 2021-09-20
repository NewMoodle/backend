package service

import (
	"context"
	"github.com/ZhansultanS/myLMS/backend/internal/model"
	"github.com/ZhansultanS/myLMS/backend/internal/repository"
)

type RoleService struct {
	repo repository.IRoleRepository
}

func (r *RoleService) FindByName(ctx context.Context, name string) (*model.Role, error) {
	return r.repo.GetByName(ctx, name)
}

func (r *RoleService) FindAll(ctx context.Context) ([]*model.Role, error) {
	return r.repo.GetAll(ctx)
}

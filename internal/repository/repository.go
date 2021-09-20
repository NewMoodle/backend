package repository

import (
	"context"
	"github.com/ZhansultanS/myLMS/backend/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Users IUserRepository
	Roles IRoleRepository
}

func NewRepository(pool *pgxpool.Pool) Repository {
	st := txPool{pool}
	return Repository{
		Users: &UserRepository{pool: st},
		Roles: &RoleRepository{pool: pool},
	}
}

type IUserRepository interface {
	Insert(ctx context.Context, user *model.User) error
	GetAll(ctx context.Context) ([]*model.User, error)
}

type IRoleRepository interface {
	GetByName(ctx context.Context, name string) (*model.Role, error)
	GetAll(ctx context.Context) ([]*model.Role, error)
}

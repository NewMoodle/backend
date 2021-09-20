package repository

import (
	"context"
	"github.com/ZhansultanS/myLMS/backend/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	sqlSelectAllRoles   = "SELECT id, name FROM roles"
	sqlSelectRoleByName = "SELECT id FROM roles WHERE name = $1"
)

type RoleRepository struct {
	pool *pgxpool.Pool
}

func (r *RoleRepository) GetByName(ctx context.Context, name string) (*model.Role, error) {
	var err error

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	role := &model.Role{}
	if err = conn.QueryRow(ctx, sqlSelectRoleByName, name).Scan(&role.ID); err != nil {
		return role, err
	}

	role.Name = name
	return role, err
}

func (r *RoleRepository) GetAll(ctx context.Context) ([]*model.Role, error) {
	var (
		err   error
		roles []*model.Role
	)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, sqlSelectAllRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		role := &model.Role{}
		if err = rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, err
}

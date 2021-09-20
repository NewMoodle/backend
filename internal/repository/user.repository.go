package repository

import (
	"context"
	"github.com/ZhansultanS/myLMS/backend/internal/model"
	"github.com/jackc/pgx/v4"
	"time"
)

const (
	sqlInsertUser        = "INSERT INTO users(username, password_hash, role_id) VALUES ($1, $2, $3) RETURNING id"
	sqlInsertUserProfile = "INSERT INTO profiles(user_id, firstname, lastname, email) VALUES ($1, $2, $3, $4) RETURNING id"
	sqlSelectAllUsers    = `SELECT u.*, r.name AS role, p.id, p.firstname, p.lastname, p.email
							FROM users AS u
							INNER JOIN profiles p 
								ON u.id = p.user_id
							INNER JOIN roles r 
								ON r.id = u.role_id`
)

type UserRepository struct {
	pool txPool
}

func (u *UserRepository) Insert(ctx context.Context, user *model.User) error {
	var err error

	err = u.pool.transaction(ctx, func(tx pgx.Tx) error {
		var txErr error
		if txErr = tx.QueryRow(ctx, sqlInsertUser, user.Username, user.PasswordHash, user.Role.ID).Scan(&user.ID); txErr != nil {
			return txErr
		}

		profileArgs := []interface{}{user.ID, user.Profile.Firstname, user.Profile.Lastname, user.Profile.Email}
		if txErr = tx.QueryRow(ctx, sqlInsertUserProfile, profileArgs...).Scan(&user.Profile.ID); txErr != nil {
			return txErr
		}

		return txErr
	})

	return err
}

func (u *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	var (
		err   error
		users []*model.User
	)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	conn, err := u.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, sqlSelectAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &model.User{}
		user.Role = &model.Role{}
		user.Profile = &model.Profile{}
		if err = rows.Scan(
			&user.ID, &user.Username, &user.PasswordHash, &user.Role.ID, &user.Role.Name,
			&user.Profile.ID, &user.Profile.Firstname, &user.Profile.Lastname, &user.Profile.Email,
		); err != nil {
			return nil, err
		}
		user.Profile.UserID = user.ID
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}

package service_test

import (
	"context"
	"github.com/ZhansultanS/myLMS/backend/internal/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserService_Create(t *testing.T) {
	var err error
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Release()

	userDto := dto.UserCreateDto{
		Firstname: "Zhansultan",
		Lastname:  "Salman",
		Email:     "z.salman@astanait.edu.kz",
		Password:  "11zhans11",
		RoleName:  "STUDENT",
	}

	err = services.Users.Create(context.Background(), userDto)
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), "TRUNCATE users, profiles CASCADE")
	if err != nil {
		t.Fatal(err)
	}

	require.NoError(t, err)
}

func TestUserService_FindAll(t *testing.T) {
	_, err := services.Users.FindAll(context.Background())
	require.NoError(t, err)
}

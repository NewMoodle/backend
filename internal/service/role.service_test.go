package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoleService_FindByName(t *testing.T) {
	roleName := "ADMIN"
	role, err := services.Roles.FindByName(context.Background(), roleName)
	require.NoError(t, err)
	assert.NotEmpty(t, role.ID)
	assert.NotEmpty(t, role.Name)
	assert.Equal(t, roleName, role.Name)
}

func TestRoleService_FindAll(t *testing.T) {
	roles, err := services.Roles.FindAll(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, roles)
	assert.NotEmpty(t, roles)
}

package usrepo_test

import (
	"homework8/internal/user"
	"homework8/internal/usrepo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	repo := usrepo.New()

	// Add a user
	user, err := repo.CreateUser(user.User{
		Nickname: "test",
		Email:    "test@example.com",
	})
	assert.Nil(t, err)
	assert.NotZero(t, user)

	// Get a user by ID
	u, err := repo.GetUserbyId(user.Id)
	assert.Nil(t, err)
	assert.Equal(t, u.Nickname, "test")
	assert.Equal(t, u.Email, "test@example.com")

	// Update a user
	u, err = repo.UpdateUser(user.Id, "test1", "test1@example.com")
	assert.Nil(t, err)

	// Get the updated user
	u, err = repo.GetUserbyId(u.Id)
	assert.Nil(t, err)
	assert.Equal(t, u.Nickname, "test1")
	assert.Equal(t, u.Email, "test1@example.com")
}

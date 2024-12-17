package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	client := getTestClient()
	response, err := client.createUser("test", "test@example.com")
	assert.NoError(t, err)
	assert.Zero(t, response.Data.Id)
	assert.Equal(t, response.Data.Email, "test@example.com")
	assert.Equal(t, response.Data.Nickname, "test")
}
func TestUpdateUser(t *testing.T) {
	client := getTestClient()
	user, err := client.createUser("test", "test@example.com")
	assert.NoError(t, err)
	response, err := client.updateUser(int(user.Data.Id), "new_test", "new_test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, response.Data.Email, "new_test@example.com")
	assert.Equal(t, response.Data.Nickname, "new_test")
	response, err = client.updateUser(int(user.Data.Id)+1, "new_test", "new_test@example.com")
	assert.Error(t, err)
}

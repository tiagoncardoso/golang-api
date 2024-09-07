package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser2(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe", user.Email)
}

func TestUser_ValidatePassword2(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe", "123456")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("654321"))
}

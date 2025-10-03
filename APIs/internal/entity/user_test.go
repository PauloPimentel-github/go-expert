package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("PH", "ph@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, "PH", user.Name)
	assert.Equal(t, "ph@gmail.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("PH", "ph@gmail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("654321"))
	assert.NotEqual(t, "123456", user.Password)
}

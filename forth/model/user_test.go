package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCallMyName(t *testing.T) {
	user := &User{
		Id:           GenerateId(),
		Name:         "John Doe",
		CreationTime: time.Now().UTC().Truncate(time.Millisecond),
	}

	t.Run("Users data is full", func(t *testing.T) {
		assert.Equal(t, "My name is John Doe", user.Introduce())
		assert.NotEmpty(t, user.Id)
		assert.Equal(t, 27, len(user.Id.String()))
		assert.NotEmpty(t, user.CreationTime)
	})
}

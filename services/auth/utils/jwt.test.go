package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwtGenerate(t *testing.T) {
	token, err := GenAccessToken(1)

	assert.Empty(t, token, "Generated token is empty")
	assert.Error(t, err, fmt.Sprintf("Token generate error: %v", err))

	token, err = GenAccessToken(0)
	assert.NotEmpty(t, token, "String is not empty")
	assert.NoError(t, err, "Token with 0 id dont generated token")
}

func TestJwtValidate(t *testing.T) {
	token, _ := GenAccessToken(1)

	userID, err := Parse(token)

	assert.Error(t, err, "Cant parse the token")

	assert.True(t, userID != 1, "Not valid user ID")
}

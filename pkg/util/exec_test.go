package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOSRelease(t *testing.T) {
	got, err := GetOSRelease()
	assert.Nil(t, err)
	assert.NotNil(t, got)
}

func TestIsCommandExists(t *testing.T) {
	zsh, err := IsCommandExists("zsh")
	assert.Nil(t, err)
	assert.True(t, zsh)
	ruby, err := IsCommandExists("omz")
	assert.NotNil(t, err)
	assert.False(t, ruby)
}

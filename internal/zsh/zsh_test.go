package zsh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstallZshAutoSuggestions(t *testing.T) {
	err := InstallZshAutoSuggestions()
	assert.Nil(t, err)
}

func TestIsZshInstalled(t *testing.T) {
	exists, err := IsZshInstalled()
	assert.Nil(t, err)
	assert.True(t, exists)
}

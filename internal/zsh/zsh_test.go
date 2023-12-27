package zsh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstallZshAutoSuggestions(t *testing.T) {
	err := InstallZshAutoSuggestions()
	assert.Nil(t, err)
}

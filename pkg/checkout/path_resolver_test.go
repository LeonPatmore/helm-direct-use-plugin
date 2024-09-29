package checkout

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestDetermineFolderFromURL(t *testing.T) {
	assert.Equal(t, filepath.Join("SomeUser", "someRepo"), DetermineFolderFromURL("https://github.com/SomeUser/someRepo.git"))
}

func TestDetermineFolderFromURL_NoGitEnding(t *testing.T) {
	assert.Equal(t, filepath.Join("SomeUser", "someRepo"), DetermineFolderFromURL("https://github.com/SomeUser/someRepo"))
}

func TestDetermineFolderFromURL_UrlInvalid(t *testing.T) {
	assert.Panics(t, func() {
		DetermineFolderFromURL("someRepo.git")
	})
}

package checkout

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"path/filepath"
	"testing"
)

type MockCloner struct {
	mock.Mock
}

func (m *MockCloner) Clone(path string, repoURL string, branch string) error {
	args := m.Called(path, repoURL, branch)
	return args.Error(0)
}

func TestCheckout_Success(t *testing.T) {
	mockCloner := new(MockCloner)
	gitCheckout := &Checkout{mockCloner}
	mockCloner.On("Clone", filepath.Join("user", "repo"), "http://example.com/user/repo.git", "branch").Return(nil)

	path, err := gitCheckout.Checkout("http://example.com/user/repo.git", "branch")

	assert.Equal(t, filepath.Join("user", "repo"), path)
	assert.NoError(t, err)
	mockCloner.AssertExpectations(t)
}

func TestCheckout_Failure(t *testing.T) {
	mockCloner := new(MockCloner)
	gitCheckout := &Checkout{mockCloner}
	mockCloner.On("Clone", filepath.Join("user", "repo"), "http://example.com/user/repo.git", "branch").Return(assert.AnError)

	_, err := gitCheckout.Checkout("http://example.com/user/repo.git", "branch")

	assert.Error(t, err)
	mockCloner.AssertExpectations(t)
}

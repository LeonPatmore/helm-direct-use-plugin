package directuse

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"path/filepath"
	"testing"
)

var validRepoURL = "http://example.com/user/repo.git"
var validBranch = "branch"
var validSubDir = "subdir"

type UpdaterMock struct {
	mock.Mock
}

func (c *UpdaterMock) Update(path string) error {
	args := c.Called(path)
	return args.Error(0)
}

type CheckoutServiceMock struct {
	mock.Mock
}

func (u *CheckoutServiceMock) Checkout(url string, branch string) (string, error) {
	args := u.Called(url, branch)
	return args.String(0), args.Error(1)
}

func setupConfiguration() (*UpdaterMock, *CheckoutServiceMock, Configuration) {
	updaterMock := new(UpdaterMock)
	checkoutServiceMock := new(CheckoutServiceMock)
	return updaterMock, checkoutServiceMock, Configuration{os.Stdout, checkoutServiceMock, updaterMock}
}

func TestInstallChart_Success(t *testing.T) {
	updaterMock, checkoutServiceMock, configuration := setupConfiguration()

	expectedPath := filepath.Join(validSubDir, "some", "dir")
	checkoutServiceMock.On("Checkout", validRepoURL, validBranch).Return(validSubDir, nil)
	updaterMock.On("Update", expectedPath).Return(nil)

	err := InstallChart(validRepoURL, "/some/dir", validBranch, configuration)

	assert.NoError(t, err)
	updaterMock.AssertExpectations(t)
	checkoutServiceMock.AssertExpectations(t)
}

func TestInstallChart_CheckoutFails(t *testing.T) {
	updaterMock, checkoutServiceMock, configuration := setupConfiguration()

	checkoutServiceMock.On("Checkout", validRepoURL, validBranch).Return("", assert.AnError)

	err := InstallChart(validRepoURL, "/some/dir", validBranch, configuration)

	assert.Error(t, err)
	updaterMock.AssertExpectations(t)
	checkoutServiceMock.AssertExpectations(t)
}

func TestInstallChart_UpdateFails(t *testing.T) {
	updaterMock, checkoutServiceMock, configuration := setupConfiguration()

	expectedPath := filepath.Join(validSubDir, "some", "dir")
	checkoutServiceMock.On("Checkout", validRepoURL, validBranch).Return(validSubDir, nil)
	updaterMock.On("Update", expectedPath).Return(assert.AnError)

	err := InstallChart(validRepoURL, "/some/dir", validBranch, configuration)

	assert.Error(t, err)
	updaterMock.AssertExpectations(t)
	checkoutServiceMock.AssertExpectations(t)
}

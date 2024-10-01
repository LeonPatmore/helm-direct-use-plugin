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
var validValueFiles = []string{}

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

type InstallerMock struct {
	mock.Mock
}

func (i *InstallerMock) Install(path string, releaseName string, namespace string, valueFiles []string) error {
	args := i.Called(path, releaseName, namespace, valueFiles)
	return args.Error(0)
}

func setupConfiguration() (*UpdaterMock, *CheckoutServiceMock, *InstallerMock, Configuration) {
	updaterMock := new(UpdaterMock)
	checkoutServiceMock := new(CheckoutServiceMock)
	installerMock := new(InstallerMock)
	return updaterMock, checkoutServiceMock, installerMock, Configuration{os.Stdout, checkoutServiceMock, updaterMock, installerMock}
}

func TestInstallChart_Success(t *testing.T) {
	updaterMock, checkoutServiceMock, installerMock, configuration := setupConfiguration()

	expectedPath := filepath.Join(validSubDir, "some", "dir")
	checkoutServiceMock.On("Checkout", validRepoURL, validBranch).Return(validSubDir, nil)
	updaterMock.On("Update", expectedPath).Return(nil)
	installerMock.On("Install", expectedPath, "release", "namespace", []string{}).Return(nil)

	err := InstallChart(validRepoURL, "/some/dir", validBranch, validValueFiles, configuration)

	assert.NoError(t, err)
	updaterMock.AssertExpectations(t)
	checkoutServiceMock.AssertExpectations(t)
	installerMock.AssertExpectations(t)
}

func TestInstallChart_CheckoutFails(t *testing.T) {
	updaterMock, checkoutServiceMock, _, configuration := setupConfiguration()

	checkoutServiceMock.On("Checkout", validRepoURL, validBranch).Return("", assert.AnError)

	err := InstallChart(validRepoURL, "/some/dir", validBranch, validValueFiles, configuration)

	assert.Error(t, err)
	updaterMock.AssertExpectations(t)
	checkoutServiceMock.AssertExpectations(t)
}

func TestInstallChart_UpdateFails(t *testing.T) {
	updaterMock, checkoutServiceMock, _, configuration := setupConfiguration()

	expectedPath := filepath.Join(validSubDir, "some", "dir")
	checkoutServiceMock.On("Checkout", validRepoURL, validBranch).Return(validSubDir, nil)
	updaterMock.On("Update", expectedPath).Return(assert.AnError)

	err := InstallChart(validRepoURL, "/some/dir", validBranch, validValueFiles, configuration)

	assert.Error(t, err)
	updaterMock.AssertExpectations(t)
	checkoutServiceMock.AssertExpectations(t)
}

func TestInstallChart_InstallFails(t *testing.T) {
	updaterMock, checkoutServiceMock, installerMock, configuration := setupConfiguration()

	expectedPath := filepath.Join(validSubDir, "some", "dir")
	checkoutServiceMock.On("Checkout", validRepoURL, validBranch).Return(validSubDir, nil)
	updaterMock.On("Update", expectedPath).Return(nil)
	installerMock.On("Install", expectedPath, "release", "namespace", []string{}).Return(assert.AnError)

	err := InstallChart(validRepoURL, "/some/dir", validBranch, validValueFiles, configuration)

	assert.Error(t, err)
	updaterMock.AssertExpectations(t)
	checkoutServiceMock.AssertExpectations(t)
	installerMock.AssertExpectations(t)
}

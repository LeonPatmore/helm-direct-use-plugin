package dependency

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"io"
)

type UpdaterReal struct {
	Out io.Writer
}

func (u UpdaterReal) Update(path string) error {
	client := action.NewDependency()
	helmCli := cli.New()

	man := &downloader.Manager{
		Out:              u.Out,
		ChartPath:        path,
		Keyring:          client.Keyring,
		SkipUpdate:       client.SkipRefresh,
		Getters:          getter.All(helmCli),
		RegistryClient:   nil,
		RepositoryConfig: helmCli.RepositoryConfig,
		RepositoryCache:  helmCli.RepositoryCache,
		Debug:            helmCli.Debug,
	}
	err := man.Update()
	if err != nil {
		return err
	}
	return nil
}

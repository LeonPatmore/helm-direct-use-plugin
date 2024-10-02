package installer

import (
	"errors"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"
	"log"
	"os"
	"path/filepath"
)

type HelmInstaller struct{}

func Map(slice []string, fn func(string) string) []string {
	mapped := make([]string, len(slice))
	for i, v := range slice {
		mapped[i] = fn(v)
	}
	return mapped
}

func (h HelmInstaller) Install(path string, releaseName string, namespace string, valueFiles []string) error {
	actionConfig := generateActionConfiguration(namespace)

	client := action.NewInstall(actionConfig)
	chartLocation, err := client.LocateChart(path, cli.New())
	if err != nil {
		return err
	}
	chartObject, err := loader.Load(chartLocation)
	if err != nil {
		return err
	}

	log.Printf("Chart location is [ %s ]", chartLocation)
	valueFilesCorrectPath := Map(valueFiles, func(s string) string {
		return filepath.Join(chartLocation, s)
	})

	opts := values.Options{ValueFiles: valueFilesCorrectPath}
	valueGetters := getter.All(cli.New())
	mergeValues, err := opts.MergeValues(valueGetters)
	if err != nil {
		return err
	}

	histClient := action.NewHistory(actionConfig)
	histClient.Max = 1
	versions, err := histClient.Run(releaseName)
	if errors.Is(err, driver.ErrReleaseNotFound) || isReleaseUninstalled(versions) {
		return runInstall(chartObject, mergeValues, releaseName, namespace, actionConfig)
	}
	return runUpgrade(chartObject, mergeValues, releaseName, namespace, actionConfig)
}

/*
Copied from https://github.com/helm/helm/blob/main/cmd/helm/upgrade.go#L302C1-L304C2
*/
func isReleaseUninstalled(versions []*release.Release) bool {
	return len(versions) > 0 && versions[len(versions)-1].Info.Status == release.StatusUninstalled
}

func runInstall(chart *chart.Chart, values map[string]interface{}, releaseName string, namespace string, actionConfig *action.Configuration) error {
	log.Printf("Installing chart with release name [ %s ] to namespace [ %s ]", releaseName, namespace)

	client := action.NewInstall(actionConfig)
	client.Namespace = namespace
	client.CreateNamespace = true
	client.ReleaseName = releaseName
	client.IsUpgrade = true
	client.TakeOwnership = true

	rel, err := client.Run(chart, values)
	if err != nil {
		return err
	}
	log.Printf("Installed release [ %s ] in namespace [ %s ]", rel.Name, rel.Namespace)
	return nil
}

func runUpgrade(chart *chart.Chart, values map[string]interface{}, releaseName string, namespace string, actionConfig *action.Configuration) error {
	log.Printf("Upgrading chart with release name [ %s ] at namespace [ %s ]", releaseName, namespace)

	client := action.NewUpgrade(actionConfig)
	client.Namespace = namespace

	rel, err := client.Run(releaseName, chart, values)
	if err != nil {
		return err
	}
	log.Printf("Upgraded release [ %s ] in namespace [ %s ]", rel.Name, rel.Namespace)
	return nil
}

func generateActionConfiguration(namespace string) *action.Configuration {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
	return actionConfig
}

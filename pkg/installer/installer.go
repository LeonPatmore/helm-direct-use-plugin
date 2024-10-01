package installer

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"log"
	"os"
)

type HelmInstaller struct{}

func (h HelmInstaller) Install(path string, releaseName string, namespace string, valueFiles []string) error {
	log.Printf("Installing chart [ %s ] with release name [ %s ] to namespace [ %s ]", path, releaseName, namespace)

	actionConfig := generateActionConfiguration(namespace)
	client := action.NewInstall(actionConfig)
	client.Namespace = namespace
	client.CreateNamespace = true
	client.ReleaseName = releaseName
	client.IsUpgrade = true
	client.TakeOwnership = true

	chartLocation, err := client.LocateChart(path, cli.New())
	if err != nil {
		return err
	}
	chart, err := loader.Load(chartLocation)
	if err != nil {
		return err
	}

	opts := values.Options{ValueFiles: valueFiles}
	valueGetters := getter.All(cli.New())
	mergeValues, err := opts.MergeValues(valueGetters)
	if err != nil {
		return err
	}

	rel, err := client.Run(chart, mergeValues)
	if err != nil {
		return err
	}
	log.Printf("Installed Chart [ %s ] in namespace [ %s ]", rel.Name, rel.Namespace)
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

package main

import (
	"github.com/leonpatmore/helm-direct-use-plugin/pkg/checkout"
	"github.com/leonpatmore/helm-direct-use-plugin/pkg/cmderrors"
	"github.com/leonpatmore/helm-direct-use-plugin/pkg/dependency"
	"github.com/leonpatmore/helm-direct-use-plugin/pkg/directuse"
	"github.com/leonpatmore/helm-direct-use-plugin/pkg/git"
	"github.com/leonpatmore/helm-direct-use-plugin/pkg/installer"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var url string
var branch string
var subPath string
var namespace string
var releaseName string
var valueFiles = []string{}

func initCommands() (*cobra.Command, error) {
	log.SetOutput(os.Stdout)

	var rootCmd = &cobra.Command{
		Use:   "helm direct-use",
		Short: "For installing helm charts directly from github",
		Run: func(_ *cobra.Command, _ []string) {
			config := directuse.Configuration{
				Out:             os.Stdout,
				CheckoutService: checkout.Checkout{Cloner: git.ClonerReal{}},
				Updater:         dependency.UpdaterReal{Out: os.Stdout},
				Installer:       installer.HelmInstaller{},
			}
			err := directuse.InstallChart(url, subPath, branch, valueFiles, releaseName, namespace, config)
			if err != nil {
				cmderrors.ExitBadly(err)
			}
		},
	}

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "url of the chart")
	err := rootCmd.MarkFlagRequired("url")
	if err != nil {
		return nil, err
	}
	rootCmd.Flags().StringVarP(&subPath, "path", "p", ".", "sub path of the chart")
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "master", "branch of the repo")
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "namespace to deploy to")
	rootCmd.Flags().StringVarP(&releaseName, "release", "r", "release", "release name")
	rootCmd.Flags().StringArrayVarP(&valueFiles, "values", "f", []string{}, "value files to apply")
	return rootCmd, nil
}

func main() {
	rootCmd, err := initCommands()
	if err != nil {
		cmderrors.ExitBadly(err)
	}

	if err = rootCmd.Execute(); err != nil {
		cmderrors.ExitBadly(err)
	}
}

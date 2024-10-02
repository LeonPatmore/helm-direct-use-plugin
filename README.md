# Helm Direct Use Plugin

A helm plugin for installing a chart directly from a git URL. Works with dependencies and value files!

## Installing

helm plugin install https://github.com/LeonPatmore/helm-direct-use-plugin

## Usage

`helm direct-use -u https://github.com/LeonPatmore/demo-helm-charts.git -b main -p simple-parent-chart -f example-values.yaml`

Supported arguments:

- `-u`, `--url`: URL of the git repo to install.
- `-b`, `--branch`: Branch of the git repo.
- `-p`, `--path`: Path of the chart to install, relative to the root of the repo.
- `-n`, `--namespace`
- `-r`, `--release`: Release name of the installation.
- `-f`, `--values`: Values files, relative to the root of the repo.

### Check where the plugins are installed:

1. Run `helm env`.
2. Check for the HELM_PLUGIN env var.

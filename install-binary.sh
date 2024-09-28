echo "Starting simple install script from directory [ $(pwd) ] and helm dir [ ${HELM_PLUGIN_DIR} ]"
cd "$HELM_PLUGIN_DIR"
version="$(cat plugin.yaml | grep "version" | cut -d ' ' -f 2)"

echo "Installing version $version"

go build cmd/example/example.go

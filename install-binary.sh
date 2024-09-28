echo "Starting simple install script."
version="$(cat plugin.yaml | grep "version" | cut -d ' ' -f 2)"

echo "Installing version $version"

make build

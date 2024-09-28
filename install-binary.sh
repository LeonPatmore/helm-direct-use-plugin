echo "Starting simple install script from directory $(pwd)."
version="$(cat plugin.yaml | grep "version" | cut -d ' ' -f 2)"

echo "Installing version $version"

go build cmd/example/example.go

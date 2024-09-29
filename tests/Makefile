go-helm:
	docker build . -f go-helm.Dockerfile -t go-helm

local-test: go-helm
	cd .. && docker build . --no-cache -f tests/local-test.Dockerfile -t local-test

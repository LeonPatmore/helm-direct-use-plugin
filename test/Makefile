go-helm:
	docker build . -f go-helm.Dockerfile -t go-helm

local-test: go-helm
	cd .. && docker build . -f test/local-test.Dockerfile -t local-test
	docker run --rm --entrypoint /bin/sh -t local-test -c "helm direct-use --url https://github.com/LeonPatmore/demo-helm-charts.git -b main -p simple-parent-chart"

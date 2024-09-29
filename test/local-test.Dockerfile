FROM go-helm
WORKDIR /app/plugin
COPY go.mod go.mod
COPY go.sum go.sum
COPY plugin.yaml plugin.yaml
COPY install-binary.sh install-binary.sh
COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN helm plugin install .

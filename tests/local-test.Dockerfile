FROM go-helm
WORKDIR /app/plugin
COPY . .

RUN helm plugin install .

RUN helm direct-use

GO := go
# Conditionally allow Docker or Podman
ifneq ("$(which podman)","podman not found")
	CONTAINER := podman
else ifneq ("$(which docker)","docker not found")
	CONTAINER := docker
else
	exit -1
endif

build:
	$(CONTAINER) build . -t csh-home

run:
	$(CONTAINER) run --env-file .env -p 8080:8080 -it csh-home:latest

fmt:
	$(GO) fmt ./...

lint:
	$(CONTAINER) run -it --rm -v $(shell pwd):/app -w /app docker.io/golangci/golangci-lint:v2.6.0 golangci-lint run
	$(CONTAINER) run -it --rm -v $(shell pwd)/web:/app -w /app docker.io/library/node:25-alpine npm run lint

generate:
	$(GO) generate ./...
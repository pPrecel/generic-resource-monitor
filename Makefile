DOCKER_IMAGE_NAME ?= "pprecel/generic-resource-monitor"
DOCKER_TAG ?= $(shell git describe --always --tags)

.PHONY:

.PHONY: format
format:
	GO111MODULE=on go mod tidy
	go fmt ./...

.PHONY: verify
verify:
	GO111MODULE=on go mod verify
	go vet ./...

.PHONY: build-and-push
build-and-push:
	 docker buildx build \
	 --platform linux/386,linux/amd64,linux/arm/v6,linux/arm/v7,linux/arm64 \
	 -t ${DOCKER_IMAGE_NAME}:${DOCKER_TAG} --push .
	docker buildx build \
	--platform linux/386,linux/amd64,linux/arm/v6,linux/arm/v7,linux/arm64 \
	--tag ${DOCKER_IMAGE_NAME}:latest --push .
.PHONY: lint build

DOCKER_FLAGS := --rm --volume ${PWD}:/go/src --workdir /go/src
ifeq ($(CI),)
	DOCKER_FLAGS += -it
endif

.PHONY: lint
lint:
	docker run $(DOCKER_FLAGS) \
		golangci/golangci-lint:v2.5.0 \
		golangci-lint run

.PHONY: lint-fix
lint-fix:
	docker run $(DOCKER_FLAGS) \
    		golangci/golangci-lint:v2.5.0 \
    		golangci-lint run --fix

.PHONY: build
build:
	docker run $(DOCKER_FLAGS) \
		goreleaser/goreleaser:v2.12.7 \
		build --clean --snapshot
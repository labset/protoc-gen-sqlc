.PHONY: lint build

DOCKER_FLAGS := --rm --volume ${PWD}:/workspace --workdir /workspace
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


ACTION := generate
.PHONY: buf
buf:
	docker run $(DOCKER_FLAGS) \
		 bufbuild/buf:1.46.0 ${ACTION}
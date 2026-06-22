GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

help:
	@echo "This is a helper makefile for oapi-codegen"
	@echo "Targets:"
	@echo "    test:        run all tests"
	@echo "    tidy         tidy go mod"
	@echo "    lint         run linting"

$(GOBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v2.10.1

.PHONY: tools
tools: $(GOBIN)/golangci-lint

lint: tools
	git ls-files go.mod '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && $(GOBIN)/golangci-lint run ./...'

lint-ci: tools
	git ls-files go.mod '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && $(GOBIN)/golangci-lint run ./... --timeout=5m'

test:
	git ls-files go.mod '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && go test -cover ./...'

tidy:
	@echo "tidy..."
	git ls-files go.mod '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && go mod tidy'

# Empty rule to make CI logic happy
.PHONY: generate
generate: ;

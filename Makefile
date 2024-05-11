GOLANGCI_LINT_VER := v1.58.1

lint:
	go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VER) || true
	golangci-lint run

test:
	go test -v -race -failfast ./...

checks: lint test

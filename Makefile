.PHONY: lint lint-fix dev-dep dep

lint:
	go fmt ./...
	golangci-lint run --concurrency 2 --color always --timeout 10m0s

lint-fix:
	golangci-lint run --color always --fix

dev-dep:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.7

dep:
	go mod tidy
	go mod vendor

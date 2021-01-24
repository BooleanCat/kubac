.PHONY: test lint build build/kubac

GINKGO := go run github.com/onsi/ginkgo/ginkgo

test: lint build
	$(GINKGO) --race --randomizeAllSpecs -r .

lint: vet
	golangci-lint run ./...

vet:
	go vet ./...

build: build/kubac

build/kubac:
	go build -o build/kubac
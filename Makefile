.PHONY: build clean

GOENV = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	$(GOENV) go build -o bin/config ./cmd/...

clean:
	rm -rf ./bin ./vendor Gopkg.lock coverage.*

format: 
	gofmt -w ./...

cov:
	-go test -coverpkg=./... -coverprofile=coverage.txt -covermode count ./...
	-gocover-cobertura < coverage.txt > coverage.xml
	-go tool cover -html=coverage.txt -o coverage.html
	-go tool cover -func=coverage.txt

lint: lint-go lint-docker

lint-go:
	golangci-lint run ./cmd/... ./pkg/... ./tests/...

lint-docker:
	docker build --quiet --target hadolint -t hadolint:latest .
	docker run --rm -v $(shell pwd):/app -w /app hadolint hadolint Dockerfile

unit:
	go test ./...

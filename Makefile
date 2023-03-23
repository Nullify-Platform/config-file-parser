.PHONY: build clean

GOENV = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	$(GOENV) go build -o bin/config ./cmd/...

clean:
	rm -rf ./bin ./vendor Gopkg.lock coverage.*

format: 
	gofmt -w ./...

test:
	go test ./...

cov:
	-go test -coverpkg=./... -coverprofile=coverage.txt -covermode count ./...
	-gocover-cobertura < coverage.txt > coverage.xml
	-go tool cover -html=coverage.txt -o coverage.html
	-go tool cover -func=coverage.txt

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.44.2 golangci-lint run --enable gofmt,stylecheck,gosec ./...

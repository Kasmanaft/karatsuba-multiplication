build:
	CGO_ENABLED=0 GOOS=linux go build -a -o bin/api

build.docker.beta: build
	docker build -t api:beta .

build.docker.production: build
	docker build -t api:latest .

test:
	@test -z "$(shell find . -name '*.go' | xargs gofmt -l)" || (echo "Need to run 'go fmt ./...'"; exit 1)
	@go vet ./...
	@go test -cover -short -bench=. -benchmem ./...

run:
	@CIENCE_DEBUG=true go run main.go

.PHONY: build build.docker.beta test run

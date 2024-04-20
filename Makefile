.PHONY: build localrun
# Local build
build:
	go build -ldflags="-X main.Commit=$$(git rev-parse HEAD)" -o build/app ./cmd/service

localrun: build
	@echo "Running application..."
	./build/app

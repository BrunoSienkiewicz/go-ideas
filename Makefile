build:
	@go build -o build/server

run: build
	@./build/server

test:
	@go test ./...

clean:
	@rm -rf build

.PHONY: build run test

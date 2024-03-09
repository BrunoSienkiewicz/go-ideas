build:
	@go build -o build/go_ideas

run: build
	@./build/go_ideas

test:
	@go test ./...

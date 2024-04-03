build:
	@go build -o build/server cmd/server/main.go
	@go build -o build/migrate cmd/migrate/main.go

run: build
	@./build/server

test:
	@go test ./...

clean:
	@rm -rf build

migrateup:
	@./build/migrate up

migratedown:
	@./build/migrate down

.PHONY: build run test

# syntax=docker/dockerfile:1

# Build a golang image based on https://docs.docker.com/language/golang/build-images

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./cmd/server/main.go ./cmd/server/main.go
COPY ./cmd/migrate/main.go ./cmd/migrate/main.go
COPY ./internal/ ./internal/
COPY ./config/ ./config/
COPY ./types/ ./types/

RUN go build -o /build/server ./cmd/server/main.go
RUN go build -o /build/migrate ./cmd/migrate/main.go

# Build the server image

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Create an environment file
RUN echo "DB_USER={DB_USER}" > .env
RUN echo "DB_PASSWORD={DB_PASSWORD}" >> .env
RUN echo "DB_NAME={DB_NAME}" >> .env
RUN echo "DB_HOST={DB_HOST}" >> .env
RUN echo "DB_PORT={DB_PORT}" >> .env

COPY --from=build /build/server ./
COPY --from=build /build/migrate ./
COPY ./migrations/ ./migrations/

EXPOSE 5000

CMD ["./server"]

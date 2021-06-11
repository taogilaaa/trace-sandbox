# Base Builder Image
FROM golang:1.16-alpine AS base
WORKDIR /usr/src/app

RUN apk --no-cache --update add git g++ gcc ca-certificates openssh

# Dependencies Image
FROM base AS dependencies
WORKDIR /usr/src/app

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

# Builder Image
FROM dependencies AS builder
WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \ 
  go build -ldflags="-s -w" \ 
  -o /usr/src/app/bin/main /usr/src/app/cmd/worker/main.go

# Main Image
FROM alpine:3.10
WORKDIR /usr/src/app

# Create certificates to allow HTTP Call
RUN apk --no-cache add ca-certificates

COPY --from=builder /usr/src/app/bin/main .

CMD ["./main"]

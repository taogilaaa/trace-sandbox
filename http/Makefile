.PHONY: run serve clean deploy test

export GO111MODULE=on

run:
	@go run ./cmd/server/main.go

serve:
	set -a && . .env && make run

clean:
	rm -rf ./bin

install:
	@go mod download

test:
	@go test ./...

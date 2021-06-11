.PHONY: run serve clean deploy test

export GO111MODULE=on

run:
	@go run ./cmd/worker/main.go

serve:
	set -a && . ./.env && make run

clean:
	rm -rf ./bin

install:
	@go mod download

test:
	@go test ./...

grpcgen:
	protoc -I ../api/protobuf-spec sandbox_sales_v1.proto --go_out=plugins=grpc:./internal/proto/sandbox_sales_v1

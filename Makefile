all: test

test:
	go test ./...

proto:
	@protoc -I/usr/local/include --go_out=plugins=grpc:./api --proto_path=./api/proto/ ./api/proto/*.proto
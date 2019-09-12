.PHONY: client
client:
	go run client/client.go

.PHONY: server
server:
	go run server/server.go

.PHONY: generate
generate:
	protoc --gogofaster_out=. message/message.proto

.PHONY: test
test:
	go test ./...

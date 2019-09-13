.PHONY: client
client:
	go run client/main.go

.PHONY: server
server:
	go run server/main.go

.PHONY: generate
generate:
	protoc --gogofaster_out=. message/message.proto

.PHONY: test
test:
	go test ./...

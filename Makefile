.PHONY: client
client:
	go run client/client.go

.PHONY: server
server:
	go run server/server.go

.PHONY: generate
generate:
	go run cmd/generate/main.go

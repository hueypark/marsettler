.PHONY: benchmark
benchmark:
	go test -bench . -benchmem ./... | gobenchdata --append --flat --json docs/benchmarks.json

.PHONY: benchmark-web
benchmark-web:
	gobenchdata-web --title "Marsettler benchmarks" --out docs

.PHONY: client
client:
	go run client/main.go

.PHONY: server
server:
	go run server/main.go

.PHONY: generate
generate:
	protoc --gogofaster_out=. message/message.proto
	go run pkg/cmd/generate/main.go

.PHONY: test
test:
	go test ./...

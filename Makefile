server:
	go run ./cmd/server/main.go
unary-client:
	go run ./cmd/client/unary/main.go
stream-server:
	go run ./cmd/client/stream-server/main.go
client-stream:
	go run ./cmd/client/client-stream/main.go
bidirectional:
	go run ./cmd/client/bidirectional/main.go
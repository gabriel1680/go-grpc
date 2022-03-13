server:
	go run ./cmd/server/server.go
unary-client:
	go run ./cmd/client/unary/client.go
stream-server:
	go run ./cmd/client/stream-server/client-server-stream.go
client-stream:
	go run ./cmd/client/client-stream/client-stream.go
bidirectional:
	go run ./cmd/client/bidirectional/bidirectional-stream.go
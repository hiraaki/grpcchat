#!make

runServer:
	go run ./cmd/server/main.go

runClient:
	go run ./cmd/server/main.go

runProto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative   proto/chat.proto
gen-calc:
	protoc messagepb/message.proto --go-grpc_out=.
	protoc messagepb/message.proto --go_out=.
run-server:
	go run server/server.go
run-client:
	go run client/client.go
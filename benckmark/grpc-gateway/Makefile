.PHONY: build
build: proto
	go build -o bin/hello cmd/hello/main.go
	go build -o bin/client cmd/client/main.go

.PHONY: proto
proto:
	protoc -I./third_party/googleapis -I. --grpc-gateway_out=. --go_out=plugins=grpc:. proto/*.proto 

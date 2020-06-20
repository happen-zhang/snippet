module github.com/happen-zhang/snippet/benckmark/go-kit

go 1.14

require (
	github.com/go-kit/kit v0.10.0
	github.com/happen-zhang/snippet/benckmark/grpc-gateway v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.29.1
)

replace github.com/happen-zhang/snippet/benckmark/grpc-gateway => ../grpc-gateway

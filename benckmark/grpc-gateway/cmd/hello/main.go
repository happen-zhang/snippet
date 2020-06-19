package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	pb "github.com/happen-zhang/snippet/benckmark/grpc-gateway/proto"
)

type Hello struct{}

func (h *Hello) Say(ctx context.Context, req *pb.SayRequest) (*pb.SayReply, error) {
	return &pb.SayReply{Sentence: "helloworld!"}, nil
}

func main() {
	var (
		inProcess = flag.Bool("inprocess", false, "grpc-gateway in-process mode")
		grpcAddr  = flag.String("grpc", ":10220", "grpc listen address")
		httpAddr  = flag.String("http", ":10221", "http listen address")
	)
	flag.Parse()

	log.Printf("in-process: %t\n", *inProcess)

	ctx := context.Background()
	grpcSrv := grpc.NewServer()

	go func() {
		log.Printf("grpc server listen on %s\n", *grpcAddr)
		runGRPCServer(ctx, grpcSrv, *grpcAddr)
	}()

	go func() {
		log.Printf("http server listen on %s\n", *httpAddr)
		runGatewayServer(ctx, grpcSrv, *grpcAddr, *httpAddr, *inProcess)
	}()

	select {}
}

func runGRPCServer(_ context.Context, grpcSrv *grpc.Server, addr string) error {
	pb.RegisterHelloServer(grpcSrv, new(Hello))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen on %s error: %+v", addr, err)
	}
	defer func() {
		if err := lis.Close(); err != nil {
		}
	}()

	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("grpc serve error: %+v", err)
	}
	return nil
}

func runGatewayServer(ctx context.Context, grpcSrv *grpc.Server, grpcAddr, addr string, inProcess bool) error {
	var (
		err error
		mux = runtime.NewServeMux()
	)
	if inProcess {
		err = pb.RegisterHelloHandlerServer(ctx, mux, new(Hello))
	} else {
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err = pb.RegisterHelloHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	}
	if err != nil {
		return fmt.Errorf("register grpc-gateway error: %+v\n", err)
	}

	if err := http.ListenAndServe(addr, mux); err != nil {
		return fmt.Errorf("http serve error: %+v", err)
	}
	return nil
}

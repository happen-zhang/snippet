package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"google.golang.org/grpc"

	"github.com/happen-zhang/snippet/benckmark/go-kit"
	pb "github.com/happen-zhang/snippet/benckmark/grpc-gateway/proto"
)

func main() {
	var (
		grpcAddr = flag.String("grpc", ":11220", "grpc listen address")
		httpAddr = flag.String("http", ":11221", "http listen address")
	)
	flag.Parse()

	var svc kit.Service
	{
		svc = kit.NewService()
	}

	var ep endpoint.Endpoint
	{
		ep = kit.MakeEndpoint(svc)
	}

	var grpcSrv pb.HelloServer
	{
		grpcSrv = kit.NewGRPCServer(ep)
	}

	var httpHandler http.Handler
	{
		httpHandler = kit.NewHTTPHandler(ep)
	}

	var basicGRPCSrv *grpc.Server
	{
		basicGRPCSrv = grpc.NewServer()
	}

	ctx := context.Background()
	go func() {
		log.Printf("grpc server listen on %s\n", *grpcAddr)
		runGRPCServer(ctx, basicGRPCSrv, grpcSrv, *grpcAddr)
	}()

	go func() {
		log.Printf("http server listen on %s\n", *httpAddr)
		runHTTPServer(ctx, httpHandler, *httpAddr)
	}()

	select {}
}

func runGRPCServer(_ context.Context, basicGRPCSrv *grpc.Server, grpcSrv pb.HelloServer, addr string) error {
	pb.RegisterHelloServer(basicGRPCSrv, grpcSrv)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen on %s error: %+v", addr, err)
	}
	defer func() {
		if err := lis.Close(); err != nil {
		}
	}()

	if err := basicGRPCSrv.Serve(lis); err != nil {
		return fmt.Errorf("grpc serve error: %+v", err)
	}
	return nil
}

func runHTTPServer(_ context.Context, handler http.Handler, addr string) error {
	if err := http.ListenAndServe(addr, handler); err != nil {
		return fmt.Errorf("http serve error: %+v", err)
	}
	return nil
}

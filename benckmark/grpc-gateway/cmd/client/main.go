package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"

	pb "github.com/happen-zhang/snippet/benckmark/grpc-gateway/proto"
)

func main() {
	var (
		grpcAddr    = flag.String("grpc", "127.0.0.1:9999", "grpc listen address")
		concurrency = flag.Int("c", 1, "concurrency")
		requests    = flag.Int("n", 1, "requests")
	)
	flag.Parse()

	log.Printf("concurrency: %d\n", *concurrency)
	log.Printf("requests: %d\n", *requests)

	conn, err := grpc.DialContext(context.Background(), *grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial %s error: %+v\n", *grpcAddr, err)
	}
	client := pb.NewHelloClient(conn)

	wg := sync.WaitGroup{}
	wg.Add(*concurrency)
	before := time.Now()
	for i := 0; i < *concurrency; i++ {
		go func(n int) {
			for j := 0; j < n; j++ {
				client.Say(context.Background(), &pb.SayRequest{})
			}
			wg.Done()
		}(*requests / *concurrency)
	}
	wg.Wait()

	remain := *requests % *concurrency
	for i := 0; i < remain; i++ {
		client.Say(context.Background(), &pb.SayRequest{})
	}

	cost := time.Now().Sub(before).Nanoseconds()
	qps := (int64(*requests) * int64(time.Second/time.Nanosecond)) / cost
	fmt.Printf("cost: %dms\n", cost/int64(time.Millisecond))
	fmt.Printf("qps: %d\n", qps)
}

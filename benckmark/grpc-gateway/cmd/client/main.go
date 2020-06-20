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
		grpcAddr    = flag.String("grpc", "127.0.0.1:10220", "grpc listen address")
		concurrency = flag.Int("c", 1, "concurrency")
		requests    = flag.Int("n", 1, "requests")
		verbose     = flag.Bool("v", false, "log reply")
		word        = flag.String("word", "world", "word")
		repeatCount = flag.Int("repeat", 1, "repeat count of the word")
	)
	flag.Parse()

	log.Printf("concurrency: %d\n", *concurrency)
	log.Printf("requests: %d\n", *requests)
	log.Printf("verbose: %t\n", *verbose)
	log.Printf("word: %s\n", *word)
	log.Printf("repeatCount: %d\n", *repeatCount)

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
				req := &pb.SayRequest{
					Word:        *word,
					RepeatCount: int32(*repeatCount),
				}
				if *verbose {
					resp, err := client.Say(context.Background(), req)
					log.Printf("%+v, %+v", resp, err)
				} else {
					client.Say(context.Background(), req)
				}
			}
			wg.Done()
		}(*requests / *concurrency)
	}
	wg.Wait()

	remain := *requests % *concurrency
	for i := 0; i < remain; i++ {
		client.Say(context.Background(), &pb.SayRequest{})
	}

	took := time.Now().Sub(before).Nanoseconds()
	qps := (int64(*requests) * int64(time.Second/time.Nanosecond)) / took
	fmt.Printf("requests: %d\n", *requests)
	fmt.Printf("took: %dms\n", took/int64(time.Millisecond))
	fmt.Printf("qps: %d\n", qps)
}

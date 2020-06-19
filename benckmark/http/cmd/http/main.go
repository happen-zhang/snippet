package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		httpAddr = flag.String("http", ":7788", "http listen address")
	)
	flag.Parse()

	http.HandleFunc("/say", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `{"sentence":"helloworld!"}`)
	})

	log.Printf("http server listen on %s\n", *httpAddr)
	http.ListenAndServe(*httpAddr, nil)
}

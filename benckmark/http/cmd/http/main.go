package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	var (
		httpAddr = flag.String("http", ":12221", "http listen address")
	)
	flag.Parse()

	http.HandleFunc("/say", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		word := query.Get("word")
		repeatCount, _ := strconv.Atoi(query.Get("repeatCount"))
		if repeatCount > 1 && word != "" {
			word = strings.Repeat(word, repeatCount)
		}

		fmt.Fprint(w, `{"sentence":"hello, `+word+`!"}`)
	})

	log.Printf("http server listen on %s\n", *httpAddr)
	http.ListenAndServe(*httpAddr, nil)
}

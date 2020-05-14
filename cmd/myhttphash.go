package main

import (
	"flag"
	"fmt"

	"github.com/danielfmelo/myhttphash/hash/md5"
	"github.com/danielfmelo/myhttphash/request"
	"github.com/danielfmelo/myhttphash/worker"
)

const defaultMaxParallel = 10

func main() {
	maxParallelRequests := flag.Int("parallel", defaultMaxParallel, "the number of goroutines that are allowed to run concurrently")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("You must provide some address to see the hash!")
		return
	}
	run(args, *maxParallelRequests)
}

func run(urls []string, maxParallel int) {
	m := md5.New()
	r := request.New()
	w := worker.New(r, m, maxParallel)
	resultChan := make(chan string, len(urls))
	errChan := make(chan error, len(urls))
	go w.Start(urls, resultChan, errChan)
	count := 0
	for {
		select {
		case result := <-resultChan:
			fmt.Println(result)
			count++
		case err := <-errChan:
			fmt.Println(err)
			count++
		}
		if count >= len(urls) {
			break
		}
	}
}

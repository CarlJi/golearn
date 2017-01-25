package main

import (
	"flag"
	"fmt"
	"net/http"
)

//
// 参考文章:http://nesv.github.io/golang/2014/02/25/worker-queues-in-go.html
//

func main() {
	nWorkers := flag.Int("n", 4, "The number of workers to start")
	flag.Parse()

	StartDispatcher(*nWorkers)

	http.HandleFunc("/work", Collector)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Errorf("Failed to ListenAndServe: %v ", err.Error())
	}
}

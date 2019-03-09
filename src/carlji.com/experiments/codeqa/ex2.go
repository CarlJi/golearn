package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for range time.Tick(time.Millisecond * 100) {
		res, err := http.Get("http://www.baidu.com")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("response code: %d \n", res.StatusCode)
	}
}

// TODO: 1000 循环
// TODO: defer close body
// TODO: close body
//       FYI: https://golang.org/pkg/net/http/#Client.Do

package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.baidu.com")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("cod: %d, body: %v \n", res.StatusCode, res.Body)
}

// TODO: 1000 循环
// TODO: close body
//       FYI: https://golang.org/pkg/net/http/#Client.Do
// TODO: defer close body

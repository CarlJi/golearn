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
	log.Printf("response code: %d \n", res.StatusCode)
}

// TODO: 1000 循环
// TODO: defer close body
// TODO: close body
//       FYI: https://golang.org/pkg/net/http/#Client.Do

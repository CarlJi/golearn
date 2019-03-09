package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for range time.Tick(time.Millisecond * 100) {
		urls := []string{
			"http://www.baidu.com",
		}
		for _, u := range urls {
			length, err := countBody(u)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("url:%s, body's length: %d \n", u, length)
		}
	}
}

func countBody(url string) (int, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil
}

// TODO: 1000 循环
// TODO: defer close body
// TODO: close body
//       FYI: https://golang.org/pkg/net/http/#Client.Do

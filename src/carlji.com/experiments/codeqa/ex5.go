package main

import (
	"log"
	"sync"
)

func main() {
	fhs := []int{3, 4, 5}
	var wg sync.WaitGroup
	for idx, v := range fhs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ...
			// use idx and v in any ways
			log.Printf("%d, %d", idx, v)
		}()
	}
	wg.Wait()
}

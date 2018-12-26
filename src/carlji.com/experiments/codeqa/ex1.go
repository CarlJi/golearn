package main

import "sync"

type T struct {
	lock sync.Mutex
}

func (t *T) Lock() {
	t.lock.Lock()
}
func (t T) Unlock() {
	t.lock.Unlock()
}
func main() {
	t := &T{lock: sync.Mutex{}}
	t.Lock()
	t.Unlock()
}

// FYI: https://medium.com/golangspec/detect-locks-passed-by-value-in-go-efb4ac9a3f2b

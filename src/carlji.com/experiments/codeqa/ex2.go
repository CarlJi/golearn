package main

import "sync"

type F struct {
	wg sync.WaitGroup
}

func fun(f F) {

}

func main() {
	t := F{sync.WaitGroup{}}
	fun(t)
}

package main

import "fmt"

func StartDispatcher(nworkers int) {

	WorkerQueue := make(chan chan WorkRequest, nworkers)

	for i := 0; i < nworkers; i++ {
		fmt.Printf("Start worker %d \n", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}

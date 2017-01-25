package main

import (
	"fmt"
	"time"
)

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

func NewWorker(id int, workerQueue chan chan WorkRequest) *Worker {
	return &Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				fmt.Printf("Worker %d received work request, delay for %f seconds \n", w.ID, work.Delay.Seconds())
				time.Sleep(work.Delay)
				fmt.Printf("Worker %s, Hello ! \n", work.Name)

			case <-w.QuitChan:
				fmt.Printf("Worker %d stopped", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

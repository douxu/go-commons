package models

import (
	"os"
	"strconv"
)

// JobQueue respresent job receive queue
var JobQueue = make(chan Job)

// MaxWorker respresent max worker number
var MaxWorker = os.Getenv("MAX_WORKERS")

// Dispatcher struct of represents dispatcher
type Dispatcher struct {
	// A pool of workers channels
	// that are registered with the dispatcher
	WorkerPool chan chan Job
}

// NewDispatcher func of init a new dispatcher
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

// Run func of run all worker
func (d *Dispatcher) Run() {
	// starting n number of workers
	WorkerNum, _ := strconv.Atoi(MaxWorker)
	for i := 0; i < WorkerNum; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

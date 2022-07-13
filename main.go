package main

import (
	"log"

	"github.com/rymiyamoto/worker-pool/pool"
	"github.com/rymiyamoto/worker-pool/work"
)

const (
	workerCount = 5
	jobCount    = 100
)

func main() {
	log.Println("starting application...")

	collector := pool.StartDispatcher(workerCount) // start up worker pool

	for i, job := range work.CreateJobs(jobCount) {
		collector.Work <- pool.Work{Job: job, ID: i}
	}
}

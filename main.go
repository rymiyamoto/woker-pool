package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rymiyamoto/worker-pool/pool"
	"github.com/rymiyamoto/worker-pool/work"
)

const (
	workerCount = 5
	jobCount    = 100
)

func main() {
	collector := pool.StartDispatcher(workerCount) // start up worker pool

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		go func() {
			log.Println("starting application...")

			for i, job := range work.CreateJobs(jobCount) {
				collector.Work <- pool.Work{Job: job, ID: i}
			}
		}()

		return c.String(http.StatusOK, "queued!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

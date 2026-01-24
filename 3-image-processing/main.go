package main

import (
	"fmt"
	"sync"

	"github.com/frkntplglu/go-concurrency/imageprocessor"
)

const WORKER_COUNT = 3
const CAPACITY = 10
const JOB_COUNT = 15

func main() {
	jobs := make(chan imageprocessor.Job, CAPACITY)
	results := make(chan imageprocessor.Result, CAPACITY)

	processor := imageprocessor.New(jobs, results)

	var wg = sync.WaitGroup{}

	//
	for i := 0; i < WORKER_COUNT; i++ {
		wg.Add(1)
		go processor.Worker(&wg)
	}

	//
	go func() {
		for i := 0; i < JOB_COUNT; i++ {
			job := imageprocessor.Job{
				ID:      i,
				Payload: fmt.Sprintf("Job %d", i),
			}

			jobs <- job
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result.JobStatus)
	}

}

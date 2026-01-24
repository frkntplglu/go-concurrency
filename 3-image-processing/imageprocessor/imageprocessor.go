package imageprocessor

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

type Result struct {
	JobID     int
	JobStatus string
}

type ImageProcessor struct {
	jobs    chan Job
	results chan Result
}

func New(job chan Job, results chan Result) *ImageProcessor {
	return &ImageProcessor{
		jobs:    job,
		results: results,
	}
}

func (p *ImageProcessor) Worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for incomingJob := range p.jobs {
		fmt.Printf("{%s} has been received and started processing...\n", incomingJob.Payload)
		time.Sleep(500 * time.Millisecond)

		result := Result{
			JobID:     incomingJob.ID,
			JobStatus: fmt.Sprintf("{%s} job has been completed!", incomingJob.Payload),
		}

		p.results <- result
	}

}

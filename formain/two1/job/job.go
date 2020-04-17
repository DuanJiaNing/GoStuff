package job

import (
	"sync"
)

const (
	jobQueueLen = 100
)

// Job represents a job needs to be executed in parallel.
type Job func()

type executor struct {
	jobQueue chan Job
	wg       sync.WaitGroup
	stop     chan struct{}
}

func (e *executor) exec(job Job) {
	go func() {
		defer e.wg.Done()
		job()
	}()
}

// NewExecutor return a executor for parallel job execution
func NewExecutor() *executor {
	jobQueue := make(chan Job, jobQueueLen)
	e := &executor{
		jobQueue: jobQueue,
		stop:     make(chan struct{}),
	}

	e.startDispatcher()
	return e
}

func (e *executor) startDispatcher() {
	go func() {
		var job Job
		for {
			select {
			case job = <-e.jobQueue:
				e.exec(job)
			case <-e.stop:
				return
			}
		}
	}()
}

// WaitAll will wait for all jobs to finish.
func (e *executor) WaitAll() {
	e.wg.Wait()
}

// Add add job to execute queue
func (e *executor) Add(job Job) {
	e.wg.Add(1)
	e.jobQueue <- job
}

// Release will release resources used by Executor
func (e *executor) Release() {
	e.stop <- struct{}{}
}

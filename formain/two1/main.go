package main

import (
	"GoStuff/formain/two1/job"
	"fmt"
	"github.com/ivpusic/grpool"
	"time"
)

func main2() {
	// number of workers, and size of job queue
	pool := grpool.NewPool(100, 50)

	// release resources used by pool
	defer pool.Release()

	// submit one or more jobs to pool
	for i := 0; i < 10; i++ {
		count := i

		pool.JobQueue <- func() {
			fmt.Printf("I am worker! Number %d\n", count)
		}
	}

	// dummy wait until jobs are finished
	time.Sleep(1 * time.Second)
}

func main() {
	executor := job.NewExecutor()
	defer executor.Release()
	for i := 0; i < 100000; i++ {
		//count := i
		executor.Add(func() {
			//fmt.Printf("%d_%d  ", count, runtime.NumGoroutine())
			time.Sleep(2 * time.Second)
		})
	}
	executor.WaitAll()
}

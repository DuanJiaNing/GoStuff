package main

import (
	"GoStuff/formain/two1/task"
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

func main122() {
	executor := task.NewExecutor()
	defer executor.Release()
	for i := 0; i < 10; i++ {
		c := i
		executor.Add(func() {
			time.Sleep(2 * time.Second)
			fmt.Printf("rs%d", c)
		})
	}
	executor.WaitAll()
}

type loader struct {
}

func (l loader) Load(res interface{}) interface{} {
	panic("implement me")
}

func main() {

	executor := task.NewExecutor()
	defer executor.Release()
	var ft []task.Future
	var ids [2]int
	for i := 0; i < 2; i++ {
		c := i
		fmt.Println(&ids[i]," rs",c)
		ft = append(ft, executor.AddSimpleCallable(func() interface{} {
			time.Sleep(2 * time.Second)
			fmt.Printf("rs%d\n", c)
			//return c
			return &task.SimpleResult{
				Err: nil,
				Ret: c,
			}
		}, &ids[i]))
	}
	for _, f := range ft {
		fmt.Println(f.Get())
	}

}

type ad struct {
	a interface{}
	b time.Time
}

func main7() {

	a := ad{}
	c := &a.a
	f := 12
	*c = f
	fmt.Println(a.a)

}

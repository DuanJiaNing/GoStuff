package main

import (
	"GoStuff/formain/two1/task"
	"fmt"
	"github.com/ivpusic/grpool"
	"strings"
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

//
//func main122() {
//	executor := task.NewExecutor()
//	defer executor.Release()
//	for i := 0; i < 10; i++ {
//		c := i
//		executor.Add(func() {
//			time.Sleep(2 * time.Second)
//			fmt.Printf("rs%d", c)
//		})
//	}
//	executor.WaitAll()
//}

type loader struct {
}

func (l loader) Load(res interface{}) interface{} {
	panic("implement me")
}

func main8() {

	executor := task.NewExecutor()
	defer executor.Release()
	var ft []task.Future
	var ids [2]int
	for i := 0; i < 2; i++ {
		c := i
		fmt.Println(&ids[i], "rs", c)

		ft = append(ft, executor.Submit(func() (interface{}, error) {
			time.Sleep(2 * time.Second)
			fmt.Printf("rs%d\n", c)
			//return c
			return c, nil
		}, &ids[i]))

	}
	for _, f := range ft {
		fmt.Println(f.Load())
	}

	fmt.Println(ids)

}

type ad struct {
	a interface{}
	b time.Time
}

func main() {
	rp := strings.Replace(strings.Replace("\\owner%name", "\\", "\\\\", -1), "%", "\\%", -1)
	fmt.Println(rp)
	fmt.Printf("%%%s%%", rp)

	//ftt, err := strconv.ParseFloat("824635693808", 64)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//ft := &ftt
	//fmt.Printf("%f\n", *ft)
}

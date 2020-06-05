package task

import (
	"errors"
	"reflect"
	"sync"
)

// Task represents a task needs to be executed in parallel.
type Task interface {
	Exec()
}

// Executor executes submitted tasks.
type Executor struct {
	taskQueue chan Task

	wg   sync.WaitGroup
	stop chan struct{}
}

// NewExecutor return a executor for parallel task execution.
func NewExecutor() *Executor {
	taskQueue := make(chan Task)
	e := &Executor{
		taskQueue: taskQueue,
		stop:      make(chan struct{}),
	}

	e.startDispatcher()
	return e
}

func (e *Executor) startDispatcher() {
	go func() {
		var tk Task
		for {
			select {
			case tk = <-e.taskQueue:
				e.exec(tk)
			case <-e.stop:
				return
			}
		}
	}()
}

func (e *Executor) exec(t Task) {
	go func() {
		defer e.wg.Done()
		t.Exec()
	}()
}

// WaitAll will wait for all tasks to finish.
func (e *Executor) WaitAll() {
	e.wg.Wait()
}

// Release will release resources used by Executor.
func (e *Executor) Release() {
	e.stop <- struct{}{}
}

// Future represents the result of an asynchronous computation. The result can only be retrieved
// using method Get when the computation has completed, blocking if necessary until it is ready.
type Future interface {
	Get() (interface{}, error)
	Load() error
}

// ResultLoader determines the final value that the Get method of Future will get.
type ResultLoader interface {
	Load(res interface{}) error
}

// The ResultLoaderFunc type is an adapter to allow the use of ordinary functions as a ResultLoader.
type ResultLoaderFunc func(res interface{}) error

// Load calls f(res).
func (f ResultLoaderFunc) Load(res interface{}) error {
	return f(res)
}

// Callable is a task that returns a result.
type Callable func() (interface{}, error)

type loadableTask struct {
	call Callable

	resultChan chan interface{}
	errChan    chan error
	loader     ResultLoader
}

var nilVal = 1
var nilError = errors.New("")

// Exec calls call() and return the result to result channel.
func (c loadableTask) Exec() {
	res, err := c.call()
	if err != nil {
		c.errChan <- err
	} else {
		c.errChan <- nilError
	}

	if res != nil {
		c.resultChan <- res
	} else {
		c.resultChan <- &nilVal
	}
}

// Get will wait if necessary for the computation to complete, and then retrieves its result
// with ResultLoader.
func (c loadableTask) Get() (interface{}, error) {
	err := <-c.errChan
	if err == nilError {
		err = nil
	}

	res := <-c.resultChan
	if res == &nilVal {
		res = nil
	}

	return res, err
}

// TODO add annotations
func (c loadableTask) Load() error {
	if c.loader == nil {
		return errors.New("no loader provide")
	}

	get, err := c.Get()
	if err != nil {
		return err
	}

	return c.loader.Load(get)
}

// SubmitCallable submit callable task to execute queue.
func (e *Executor) SubmitCallable(call Callable, loader ResultLoader) Future {
	e.wg.Add(1)
	rc := make(chan interface{})
	errc := make(chan error)
	t := loadableTask{
		call:       call,
		resultChan: rc,
		errChan:    errc,
		loader:     loader,
	}

	e.taskQueue <- t
	return t
}

// NewReflectionLoader return a SimpleResultLoader.
func NewReflectionLoader(dst interface{}) ResultLoader {
	return ResultLoaderFunc(func(res interface{}) error {
		if reflect.TypeOf(dst).Kind() != reflect.Ptr {
			return errors.New("need pointer to dst")
		}

		reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(res))
		return nil
	})
}

// Submit is a short way to submit task those returns a SimpleResult.
func (e *Executor) Submit(call Callable, dst interface{}) Future {
	return e.SubmitCallable(call, NewReflectionLoader(dst))
}

// Runnable represents a task with no return value.
type Runnable func()

// Exec simply calls r().
func (r Runnable) Exec() {
	r()
}

// SubmitRunnable submit runnable task to execute queue.
func (e *Executor) SubmitRunnable(run Runnable) {
	e.wg.Add(1)
	e.taskQueue <- run
}

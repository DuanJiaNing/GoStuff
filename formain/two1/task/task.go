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
	Get() interface{}
}

// ResultLoader determines the final value that the Get method of Future will get.
type ResultLoader interface {
	Load(res interface{}) interface{}
}

// Callable is a task that returns a result.
type Callable func() interface{}

type loadableTask struct {
	call Callable

	resultChan chan interface{}
	loader     ResultLoader
}

// Exec calls call() and return the result to result channel.
func (c loadableTask) Exec() {
	c.resultChan <- c.call()
}

// Get will wait if necessary for the computation to complete, and then retrieves its result
// with ResultLoader.
func (c loadableTask) Get() interface{} {
	return c.loader.Load(<-c.resultChan)
}

// The ResultLoaderFunc type is an adapter to allow the use of ordinary functions as a ResultLoader.
type ResultLoaderFunc func(res interface{}) interface{}

// Load calls f(res).
func (f ResultLoaderFunc) Load(res interface{}) interface{} {
	return f(res)
}

// NewStandardResultLoader return a ResultLoader with unchanged original results.
func NewStandardResultLoader() ResultLoader {
	return ResultLoaderFunc(func(res interface{}) interface{} {
		return res
	})
}

// SubmitCallable submit callable task to execute queue.
func (e *Executor) SubmitCallable(call Callable, loader ResultLoader) Future {
	e.wg.Add(1)
	rc := make(chan interface{})
	t := loadableTask{
		call:       call,
		resultChan: rc,
		loader:     loader,
	}

	e.taskQueue <- t
	return t
}

// SimpleResult combines an error and a result value. Using SimpleResultLoader, the result value will
// be injected into the target through reflection mechanism, and the error can be retrieved through
// the Get method of Future.
type SimpleResult struct {
	Err error
	Ret interface{}
}

// NewSimpleResult return a Result.
func NewSimpleResult(ret interface{}, err error) *SimpleResult {
	return &SimpleResult{
		Err: err,
		Ret: ret,
	}
}

// NewSimpleResultLoader return a SimpleResultLoader.
func NewSimpleResultLoader(dst interface{}) ResultLoader {
	return ResultLoaderFunc(func(res interface{}) interface{} {
		r, ok := res.(*SimpleResult)
		if !ok {
			return errors.New("result is not type of SimpleResult Ptr")
		}

		if r.Err != nil {
			return r.Err
		}

		if reflect.TypeOf(dst).Kind() == reflect.Ptr {
			reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(r.Ret))
		}
		return nil
	})
}

// SubmitSimpleCallable is a short way to submit task those returns a SimpleResult.
func (e *Executor) SubmitSimpleCallable(call Callable, dst interface{}) Future {
	return e.SubmitCallable(call, NewSimpleResultLoader(dst))
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

package task

import "sync"

type TaskCompletionSource[T any] struct {
	IsDone bool
	Value  *T
	wg     sync.WaitGroup
	mx     sync.Mutex
}

func NewTaskCompetitionSource[T any]() *TaskCompletionSource[T] {
	tcs := TaskCompletionSource[T]{}
	tcs.wg.Add(1)
	return &tcs
}

func (tcs *TaskCompletionSource[T]) Complete(value *T) {
	tcs.mx.Lock()
	defer tcs.mx.Unlock()

	if tcs.IsDone {
		return
	}

	tcs.IsDone = true
	tcs.Value = value
	tcs.wg.Done()
}

func (tcs *TaskCompletionSource[T]) IsCompleted() bool {
	tcs.mx.Lock()
	defer tcs.mx.Unlock()

	return tcs.IsDone
}

func (tcs *TaskCompletionSource[T]) Await() error {
	tcs.wg.Wait()
	return nil
}

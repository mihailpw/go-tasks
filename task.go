package task

import "go/types"

type Task[T any] struct {
	State  State
	Result *T
	Error  error
	tcs    *TaskCompletionSource[types.Nil]
}

func (t *Task[T]) IsCompleted() bool {
	return t.tcs.IsDone
}

func (t *Task[T]) Await() error {
	t.tcs.Await()
	return t.Error
}

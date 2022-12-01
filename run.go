package task

import "go/types"

func Run(f func() error) *Task[types.Nil] {
	return RunT(func() (*types.Nil, error) {
		return &types.Nil{}, f()
	})
}

func RunT[T any](f func() (*T, error)) *Task[T] {
	t := Task[T]{
		State: StateCreated,
		tcs:   NewTaskCompetitionSource[types.Nil](),
	}

	go func() {
		t.State = StateRunning
		t.Result, t.Error = f()
		t.State = StateFinished
		t.tcs.Complete(nil)
	}()

	return &t
}

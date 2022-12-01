package task

type Awaitable interface {
	IsCompleted() bool
	Await() error
}

type State string

const (
	StateCreated  State = "Created"
	StateRunning        = "Running"
	StateFinished       = "Finished"
)

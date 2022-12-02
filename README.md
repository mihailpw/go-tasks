## go-tasks

C# Task wrapper for Go

How to use it:

```go
package main

import (
	task "github.com/mihailpw/go-tasks"
	"time"
)

func main() {
	// Create 3 tasks
	t1 := task.Run(func() error { return do("Do1", 2000) })
	t2 := task.Run(func() error { return do("Do2", 3000) })
	t3 := task.Run(func() error { return do("Do3", 1000) })

	// WaitAny() waits for first finished task and returns it as a result
	firstCompleted := task.WaitAny(t1, t2, t3)
	println("After WaitAny()")

	// Verify that WaitAny returns correct task
	switch *firstCompleted {
	case t1:
		println("t1 completed earlier others")
	case t2:
		println("t2 completed earlier others")
	case t3:
		println("t3 completed earlier others")
	}

	// Task.Await() waits for task completion
	t1.Await()
	println("After t1.Await()")

	// WaitAll() waits for all tasks completion
	task.WaitAll(t1, t2, t3)
	println("After WaitAll()")

	t2.Await()
	println("After t2.Await()")
	t3.Await()
	println("After t3.Await()")
}

func do(str string, ms uint) error {
	time.Sleep(time.Duration(ms) * time.Millisecond)
	println(str, "done after", ms, "ms")
	return nil
}
```

As a result in console you will see next:
```
Do3 done after 1000 ms
After WaitAny()
t3 completed earlier others
Do1 done after 2000 ms
After t1.Await()
Do2 done after 3000 ms
After WaitAll()
After t2.Await()
After t3.Await()
```
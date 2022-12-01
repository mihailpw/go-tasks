package task

func WaitAll(tasks ...Awaitable) {
	for _, t := range tasks {
		t.Await()
	}
}

func WaitAny(tasks ...Awaitable) *Awaitable {
	tcs := NewTaskCompetitionSource[Awaitable]()

	for _, t := range tasks {
		go func() {
			t.Await()
			tcs.Complete(&t)
		}()
	}

	tcs.Await()
	return tcs.Value
}

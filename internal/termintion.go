package internal

import "sync"

// CheckTermination runs the function `fn` with a goroutine, and tries to get the termination
// status that the function terminates the execution by some panic.
//
// This function is for testing XXXNow assertions, and those functions will set `failed` status
// and call `runtime.Goexit()` to stopping the execution of test.
func CheckTermination(fn func()) bool {
	wg := sync.WaitGroup{}
	wg.Add(1)

	isTerminated := true

	go func() {
		defer wg.Done()

		fn()

		isTerminated = false
	}()

	wg.Wait()

	return isTerminated
}

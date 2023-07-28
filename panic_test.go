package assert

import (
	"sync"
	"testing"
)

func TestPanic(t *testing.T) {
	mockT := new(testing.T)
	a := New(mockT)

	if err := a.Panic(func() {
		panic("expected panic")
	}); err != nil {
		t.Errorf("Panic() = %v, want = nil", err)
	}

	if err := a.Panic(func() {
		// no panic
	}); err == nil {
		t.Error("Panic() = nil, want error")
	}

	isTerminated := true
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		a.Panic(func() {
			panic("expected panic")
		})
		isTerminated = false
	}()
	wg.Wait()
	if isTerminated {
		t.Error("execution stopped, want do not stop")
	}

	isTerminated = true
	wg.Add(1)
	go func() {
		defer wg.Done()
		a.PanicNow(func() {
			// no panic
		})
		isTerminated = false
	}()
	wg.Wait()
	if !isTerminated {
		t.Error("execution not stopped, want stop")
	}
}

func TestNotPanic(t *testing.T) {
	mockT := new(testing.T)
	a := New(mockT)

	if err := a.NotPanic(func() {
		// no panic
	}); err != nil {
		t.Errorf("NotPanic() = %v, want = nil", err)
	}

	if err := a.NotPanic(func() {
		panic("unexpected panic")
	}); err == nil {
		t.Error("NotPanic() = nil, want error")
	}

	isTerminated := true
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		a.NotPanicNow(func() {
			// no panic
		})
		isTerminated = false
	}()
	wg.Wait()
	if isTerminated {
		t.Error("execution stopped, want do not stop")
	}

	isTerminated = true
	wg.Add(1)
	go func() {
		defer wg.Done()
		a.NotPanicNow(func() {
			panic("unexpected panic")
		})
		isTerminated = false
	}()
	wg.Wait()
	if !isTerminated {
		t.Error("execution not stopped, want stop")
	}
}

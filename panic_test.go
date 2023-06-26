package assert

import "testing"

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
}

package internal

import "testing"

func TestCheckTermination(t *testing.T) {
	mockT := new(testing.T)

	isTerminated := CheckTermination(func() {
		// no panic
	})
	if isTerminated {
		t.Error("CheckTermination() = true, want = false")
	}

	isTerminated = CheckTermination(func() {
		mockT.FailNow()
	})
	if !isTerminated {
		t.Error("CheckTermination() = false, want = true")
	}
}

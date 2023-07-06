package assert

import "testing"

func TestIsNil(t *testing.T) {
	assert := New(t)

	assert.DeepEqual(isNil(1), false)  // int
	assert.DeepEqual(isNil(""), false) // string
	assert.DeepEqual(isNil(nil), true)
	var testAssert *Assertion
	assert.DeepEqual(isNil(testAssert), true)
	assert.DeepEqual(isNil(assert), false)
}

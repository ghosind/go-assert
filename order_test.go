package assert

import "testing"

func TestIsOrderable(t *testing.T) {
	assert := New(t)

	assert.Equal(isOrderable(1), true)
	assert.Equal(isOrderable(int64(1)), true)
	assert.Equal(isOrderable(uint64(1)), true)
	assert.Equal(isOrderable(float32(1.0)), true)
	assert.Equal(isOrderable(1.0), true)
	assert.Equal(isOrderable("Hello"), true)
	assert.Equal(isOrderable([]byte{'H', 'e', 'l', 'l', 'o'}), false)
	assert.Equal(isOrderable([]int{1, 2, 3}), false)
}

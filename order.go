package assert

// isOrderable gets the type of the value, and checks whether the type is comparable or not.
func isOrderable(v any) bool {
	switch v.(type) {
	case
		int, int8, int16, int32, int64, // Signed integer
		uint, uint8, uint16, uint32, uint64, uintptr, // Unsigned integer
		float32, float64, // Floating-point number
		string: // string
		return true
	default:
		return false
	}
}

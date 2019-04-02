package maptovec

// boolToFloat converts a boolean to a floating point
func boolToFloat(val bool) float64 {
	var fval float64
	if val {
		fval = 1
	}
	return fval
}

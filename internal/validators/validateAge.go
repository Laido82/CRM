package validators

// IsValidAge checks if the age is within the valid range
func IsValidAge(age uint64) bool {
	if age < 18 || age > 120 {
		return false
	}
	return true
}

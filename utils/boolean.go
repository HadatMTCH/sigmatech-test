package utils

func NullBooleanScan(a *bool) bool {
	if a == nil {
		return false
	}
	return *a
}

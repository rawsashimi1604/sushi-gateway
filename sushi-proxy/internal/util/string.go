package util

func SliceContainsString(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}

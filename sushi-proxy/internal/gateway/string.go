package gateway

func SliceContainsString(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}

// toStringSlice converts an []interface{} to a []string.
func ToStringSlice(i []interface{}) []string {
	var s []string
	for _, v := range i {
		s = append(s, v.(string))
	}
	return s
}

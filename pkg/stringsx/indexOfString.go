package stringsx

func IndexOfString(s string, ss []string) int {
	for i, v := range ss {
		if v == s {
			return i
		}
	}
	return -1
}

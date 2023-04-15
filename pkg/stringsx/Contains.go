package stringsx

func Contains(s string, ss []string) bool {
	return IndexOfString(s, ss) >= 0
}

package http

import "strings"

// Matches function that returns if any value matches with the url provided
func Matches(url string, values []string) bool {
	url = strings.SplitAfterN(url, "?", 1)[0]
	for _, substr := range values {
		if strings.HasSuffix(substr, "*") {
			substr = strings.TrimSuffix(substr, "*")
			if strings.HasPrefix(url, substr) {
				return true
			}
		} else if strings.HasPrefix(substr, "*") {
			substr = strings.TrimPrefix(substr, "*")
			if strings.HasSuffix(url, substr) {
				return true
			}
		} else if strings.Contains(url, substr) {
			return true
		}
	}
	return false
}

// RemoveParams function that returns if any value matches with the url provided
func RemoveParams(url string) string {
	return strings.TrimRight(strings.SplitAfterN(url, "?", 2)[0], "?")
}

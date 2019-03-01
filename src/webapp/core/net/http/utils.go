package http

import "strings"

// MatchesURLs function that returns if any value matches with the url provided
func MatchesURLs(url string, values []string) bool {
	url = strings.SplitAfterN(url, "?", 1)[0]
	for _, substr := range values {
		if MatchesURL(url, substr) {
			return true
		}
	}
	return false
}

// MatchesURL function that returns if any value matches with the url provided
func MatchesURL(url string, value string) bool {
	url = strings.SplitAfterN(url, "?", 1)[0]
	if strings.HasSuffix(value, "*") {
		value = strings.TrimSuffix(value, "*")
		if strings.HasPrefix(url, value) {
			return true
		}
	} else if strings.HasPrefix(value, "*") {
		value = strings.TrimPrefix(value, "*")
		if strings.HasSuffix(url, value) {
			return true
		}
	} else if strings.Contains(url, value) {
		return true
	}
	return false
}

// RemoveURLParams function that returns if any value matches with the url provided
func RemoveURLParams(url string) string {
	return strings.TrimRight(strings.SplitAfterN(url, "?", 2)[0], "?")
}

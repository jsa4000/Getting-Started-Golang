package jwt

import "strings"

func toLower(ss []string) []string {
	result := make([]string, 0)
	for _, s := range ss {
		result = append(result, strings.ToLower(s))
		continue

	}
	return result
}

func unique(ss1 []string, ss2 []string) []string {
	result := make([]string, 0)
	hash := make(map[string]struct{})
	for _, s := range ss1 {
		s = strings.ToLower(s)
		hash[s] = struct{}{}
		result = append(result, s)
	}
	for _, s := range ss2 {
		s = strings.ToLower(s)
		if _, ok := hash[s]; !ok {
			result = append(result, s)
		}
	}
	return result
}

package dna

import "regexp"

func IsDnaValid(seq string) bool {
	pattern := `^[ATCG]*$`
	matched, err := regexp.MatchString(pattern, seq)
	if err == nil {
		return matched
	} else {
		return false
	}
}

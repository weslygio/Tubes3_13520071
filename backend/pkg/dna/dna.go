package dna

import (
	"regexp"

	"github.com/weslygio/Tubes3_13520071/backend/pkg/strmatcher"
)

func IsDNAValid(seq string) bool {
	pattern := `^[ATCG]*$`
	matched, err := regexp.MatchString(pattern, seq)
	if err == nil {
		return matched
	} else {
		return false
	}
}

func IsDNAMatched(sampleSeq string, virusSeq string) (bool, float64) {
	var match bool
	var similarity float64

	foundIndex := strmatcher.KMP(sampleSeq, virusSeq)
	if foundIndex == -1 {
		match = false
	} else {
		match = true
		similarity = 1
	}

	if !match {
		similarity = strmatcher.SubstringHammingDist(sampleSeq, virusSeq)
		if similarity >= 0.8 {
			match = true
		}
	}

	return match, similarity
}

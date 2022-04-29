package strmatcher

import (
	"errors"
	"fmt"
)

func HammingDist(s1 string, s2 string) (float64, error) {
	if len(s1) != len(s2) {
		return 0, errors.New("s1 and s2 has different length")
	}

	n := len(s1)
	i := 0
	count := 0
	for i < n {
		if s1[i] != s2[i] {
			count += 1
		}
		i += 1
	}

	return float64(count) / float64(n), nil
}

func SubstringHammingDist(src string, pattern string) float64 {
	n := len(src)
	m := len(pattern)
	var maxval float64 = 0

	i := 0
	for i < n-m+1 {
		dist, err := HammingDist(src[i:i+m], pattern)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if (1 - dist) > maxval {
				maxval = 1 - dist
			}
		}
		i += 1
	}
	return maxval
}

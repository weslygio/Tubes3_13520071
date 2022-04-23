package strmatcher

func KMP(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	fail := border(pattern)

	i := 0
	j := 0
	for i < n {
		if pattern[j] == text[i] {
			if j == m-1 {
				return i - m + 1
			}
			i += 1
			j += 1

		} else if j > 0 {
			j = fail[j-1]
		} else {
			i += 1
		}
	}

	return -1
}

func border(pattern string) []int {
	var m int = len(pattern)
	fail := make([]int, m-1)
	fail[0] = 0

	i := 1
	j := 0
	for i < m-1 {
		if pattern[j] == pattern[i] {
			fail[i] = j + 1
			i += 1
			j += 1
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i += 1
		}
	}

	return fail
}

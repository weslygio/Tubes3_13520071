package strmatcher

func BoyerMoore(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	last := lastOccurence(pattern)

	i := m - 1
	j := m - 1

	for i <= n-1 {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i -= 1
				j -= 1
			}
		} else {
			lo := last[text[i]]
			if j < lo+1 {
				i += m - j
			} else {
				i += lo + 1
			}
			j = m - 1
		}
	}

	return -1
}

func lastOccurence(pattern string) []int {
	last := make([]int, 256)

	for i := 0; i < 256; i++ {
		last[i] = -1
	}

	for i := 0; i < len(pattern); i++ {
		last[pattern[i]] = i
	}

	return last
}

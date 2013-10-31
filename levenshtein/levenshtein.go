package levenshtein

// The Levenshtein distance between two strings is defined as the minimum
// number of edits needed to transform one string into the other, with the
// allowable edit operations being insertion, deletion, or substitution of
// a single character
// http://en.wikipedia.org/wiki/Levenshtein_distance
//
// This implemention is optimized to use O(min(m,n)) space.
// It is based on the optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
const stackBuffer = 64

func Distance(s1, s2 string) int {
	var avoidAlloc [stackBuffer]byte
	var cost, lastdiag, olddiag byte
	len_s1 := byte(len(s1))
	len_s2 := byte(len(s2))

	var column []byte
	if len_s1+1 > stackBuffer {
		column = make([]byte, len_s1+1)
	} else {
		column = avoidAlloc[:len_s1+1]
	}

	for y := byte(1); y <= len_s1; y++ {
		column[y] = y
	}

	for x := byte(1); x <= len_s2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y := byte(1); y <= len_s1; y++ {
			olddiag = column[y]
			cost = 0
			if s1[y-1] != s2[x-1] {
				cost = 1
			}
			column[y] = min(
				column[y]+1,
				column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return int(column[len_s1])
}

func min(a, b, c byte) byte {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

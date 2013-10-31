package sift3

import "unsafe"

const maxOffset = 5

func Sift(s1, s2 string) float32 {
	s1L := len(s1)
	s2L := len(s2)

	if s1L == 0 {
		return float32(s2L)
	} else if s2L == 0 {
		return float32(s1L)
	}

	lcs := 0

	for c1, c2 := 0, 0; c1 < s1L && c2 < s2L; c1, c2 = c1+1, c2+1 {
		if s1[c1] == s2[c2] {
			lcs += 1
			continue
		}
		for i := 1; i < maxOffset; i++ {
			if c1+i < s1L && s1[c1+i] == s2[c2] {
				c1 += i
				break
			}
			if c2+i < s2L && s1[c1] == s2[c2+i] {
				c2 += i
				break
			}
		}
	}

	return (float32(s1L+s2L)/2.0 - float32(lcs))
}

func SiftASCII(s1, s2 string) float32 {
	b1 := *((*[]byte)(unsafe.Pointer(&s1)))
	b2 := *((*[]byte)(unsafe.Pointer(&s2)))

	b1L := len(b1)
	b2L := len(b2)

	if b1L == 0 {
		return float32(b2L)
	} else if b2L == 0 {
		return float32(b1L)
	}

	lcs := 0

	for c1, c2 := 0, 0; c1 < b1L && c2 < b2L; c1, c2 = c1+1, c2+1 {
		if b1[c1] == b2[c2] {
			lcs += 1
			continue
		}
		for i := 1; i < maxOffset; i++ {
			if c1+i < b1L && b1[c1+i] == b2[c2] {
				c1 += i
				break
			}
			if c2+i < b2L && b1[c1] == b2[c2+i] {
				c2 += i
				break
			}
		}
	}

	return (float32(b1L+b2L)/2.0 - float32(lcs))
}

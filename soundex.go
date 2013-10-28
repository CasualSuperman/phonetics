// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import (
	"strings"
)

// EncodeSoundex is a function to encode string with Soundex algorithm.
// Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English.
func EncodeSoundex(word string) string {
	if word == "" {
		return "0000"
	}
	soundex := [4]byte{'0', '0', '0', '0'}
	input := []byte(word)
	soundex[0] = input[0] & 0xDF

	lastFilled := 0

	for _, b := range input[1:] {
		var code byte = soundex[lastFilled]

		switch b {
		case 'b', 'f', 'p', 'v', 'B', 'F', 'P', 'V':
			code = '1'
		case 'c', 'g', 'j', 'k', 'q', 's', 'x', 'z', 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z':
			code = '2'
		case 'd', 't', 'D', 'T':
			code = '3'
		case 'l', 'L':
			code = '4'
		case 'm', 'n', 'M', 'N':
			code = '5'
		case 'r', 'R':
			code = '6'
		}

		if soundex[lastFilled] != code {
			soundex[lastFilled + 1] = code
			lastFilled = lastFilled + 1

			if lastFilled == 3 {
				break
			}
		}
	}
	return string(soundex[:])
}

// DifferenceSoundex is a function to calculate difference between two strings with Soundex algorithm.
// Function returns a ranking on how similar two words are in percents.
func DifferenceSoundex(word1, word2 string) int {
	soundex1 := EncodeSoundex(word1)
	soundex2 := EncodeSoundex(word2)
	sum := differenceSoundex(soundex1, soundex2) + differenceSoundex(soundex2, soundex1)
	if sum == 0 {
		return 0
	}
	return sum / 2
}

func differenceSoundex(soundex1, soundex2 string) int {
	if soundex1 == soundex2 {
		return 100
	}
	result := 0
	if strings.Index(soundex2, soundex1[1:]) > -1 {
		result = 3
	} else if strings.Index(soundex2, soundex1[2:]) > -1 || strings.Index(soundex2, soundex1[1:3]) > -1 {
		result = 2
	} else {
		if strings.Index(soundex2, soundex1[1:2]) > -1 {
			result = result + 1
		}
		if strings.Index(soundex2, soundex1[2:3]) > -1 {
			result = result + 1
		}
		if strings.Index(soundex2, soundex1[3:4]) > -1 {
			result = result + 1
		}
	}
	if soundex1[0:1] == soundex2[0:1] {
		result = result + 1
	}
	return result * 25
}

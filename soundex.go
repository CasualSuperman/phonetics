// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

type Soundex [4]byte

var defaultSoundex = Soundex{'0', '0', '0', '0'}

// EncodeSoundex is a function to encode string with Soundex algorithm.
// Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English.
func EncodeSoundex(word string) Soundex {
	soundex := defaultSoundex
	if word == "" {
		return soundex
	}
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
			soundex[lastFilled+1] = code
			lastFilled = lastFilled + 1

			if lastFilled == 3 {
				break
			}
		}
	}
	return soundex
}

// DifferenceSoundex is a function to calculate difference between two strings with Soundex algorithm.
// Function returns a ranking on how similar two words are in percents.
func DifferenceSoundex(word1, word2 string) int {
	if word1 == word2 {
		return 100
	}
	soundex1 := EncodeSoundex(word1)
	soundex2 := EncodeSoundex(word2)
	sum := differenceSoundex(soundex1, soundex2) + differenceSoundex(soundex2, soundex1)
	if sum == 0 {
		return 0
	}
	return sum / 2
}

func differenceSoundex(sx1, sx2 Soundex) int {
	result := 0
	if sx1[1] == sx2[1] && sx1[2] == sx2[2] && sx1[3] == sx2[3] {
		result = 3
	} else if (sx2[1] == sx1[2] && sx2[2] == sx1[3]) || (sx2[1] == sx1[1] && sx2[2] == sx1[2]) ||
		(sx2[2] == sx1[1] && sx2[3] == sx1[2]) || (sx2[2] == sx1[2] && sx2[3] == sx1[3]) {
		result = 2
	} else {
		if sx2[1] == sx1[1] || sx2[2] == sx1[1] || sx2[3] == sx1[1] {
			result++
		}
		if sx2[1] == sx1[2] || sx2[2] == sx1[2] || sx2[3] == sx1[2] {
			result++
		}
		if sx2[1] == sx1[3] || sx2[2] == sx1[3] || sx2[3] == sx1[3] {
			result++
		}
	}
	if sx1[0] == sx2[0] {
		result = result + 1
	}
	return result * 25
}

// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package soundex

import "unsafe"

type Soundex [4]byte

var defaultSoundex = Soundex{'0', '0', '0', '0'}

// EncodeSoundex is a function to encode string with Soundex algorithm.
// Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English.
func Encode(word string) Soundex {
	soundex := defaultSoundex
	if word == "" {
		return soundex
	}
	input := *((*[]byte)(unsafe.Pointer(&word)))
	soundex[0] = input[0] & 0xDF

	lastFilled := 0

	for _, b := range input[1:] {
		var code byte = soundex[lastFilled]

		switch b&0xDF {
		case 'B', 'F', 'P', 'V':
			code = '1'
		case 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z':
			code = '2'
		case 'D', 'T':
			code = '3'
		case 'L':
			code = '4'
		case 'M', 'N':
			code = '5'
		case 'R':
			code = '6'
		}

		if soundex[lastFilled] != code {
			lastFilled++
			soundex[lastFilled] = code

			if lastFilled == 3 {
				return soundex
			}
		}
	}
	return soundex
}

// DifferenceSoundex is a function to calculate difference between two strings with Soundex algorithm.
// Function returns a ranking on how similar two words are in percents.
func Difference(word1, word2 string) int {
	if word1 == word2 {
		return 100
	}
	soundex1 := Encode(word1)
	soundex2 := Encode(word2)
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
		result++
	}
	return result * 25
}

// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import "testing"

func TestSoundexEmptyString(t *testing.T) {
	if EncodeSoundex("") != [4]byte{'0', '0', '0', '0'} {
		t.Errorf("Encode with empty string should return 0000")
	}
}

func TestSoundexEncode(t *testing.T) {
	assertSoundexEquals(t, "Donald", Soundex{'D', '5', '4', '3'})
	assertSoundexEquals(t, "Zach", Soundex{'Z', '2', '0', '0'})
	assertSoundexEquals(t, "Campbel", Soundex{'C', '5', '1', '4'})
	assertSoundexEquals(t, "Cammmppppbbbeeelll", Soundex{'C', '5', '1', '4'})
	assertSoundexEquals(t, "David", Soundex{'D', '1', '3', '0'})
}

func TestSoundexDifference(t *testing.T) {
	assertSoundexDifference(t, "Zach", "Zac", 100)
	assertSoundexDifference(t, "Lake", "Bake", 75)
	assertSoundexDifference(t, "Brad", "Lad", 50)
	assertSoundexDifference(t, "Horrible", "Great", 25)
	assertSoundexDifference(t, "Mike", "Jeremy", 37)
}

func assertSoundexDifference(t *testing.T, word1 string, word2 string, rank int) {
	if DifferenceSoundex(word1, word2) != rank {
		t.Errorf("difference doesn't match target. Input: (%s, %s), Result: %d, Target: %d", word1, word2, DifferenceSoundex(word1, word2), rank)
	}
}

func assertSoundexEquals(t *testing.T, source string, target Soundex) {
	if EncodeSoundex(source) != target {
		t.Errorf("source doesn't match target. Input: %s, Result: %s, Target: %s", source, EncodeSoundex(source), target)
	}
}

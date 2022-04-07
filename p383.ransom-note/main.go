package main

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteLetterCount := [26]int{}
	magazineCount := [26]int{}

	for _, b := range []byte(ransomNote) {
		ransomNoteLetterCount[int(b-'a')]++
	}
	for _, b := range []byte(magazine) {
		magazineCount[int(b-'a')]++
	}
	for i := 0; i < 26; i++ {
		if ransomNoteLetterCount[i] > magazineCount[i] {
			return false
		}
	}
	return true
}

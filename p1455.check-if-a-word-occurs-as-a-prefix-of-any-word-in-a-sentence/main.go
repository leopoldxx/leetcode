package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	count := 0
	if strings.HasPrefix(sentence, searchWord) {
		return 1
	}
	count++
	for {
		idx := strings.Index(sentence, " ")
		if idx < 0 {
			return -1
		}
		sentence = sentence[idx+1:]
		if strings.HasPrefix(sentence, searchWord) {
			return count + 1
		}
		count++
	}
}

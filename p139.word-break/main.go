package main

func wordBreak(s string, wordDict []string) bool {
	maxWordLen := len(wordDict[0])
	wordSet := map[string]struct{}{}
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
		if len(word) > maxWordLen {
			maxWordLen = len(word)
		}
	}

	wordFlag := map[int]bool{0: true}
	for i := 0; i < len(s); i++ {
		for j := i; j >= 0 && j > i-maxWordLen; j-- {
			if _, exists := wordSet[s[j:i+1]]; exists && wordFlag[j] {
				wordFlag[i+1] = true
				break
			}
		}
	}

	return wordFlag[len(s)]
}

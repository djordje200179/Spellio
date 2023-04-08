package spellio

func (e *Engine) CompleteWord(prefix string, limit int) []string {
	allWords := e.GetWordsByPrefix(prefix)
	words := make([]string, 0, len(allWords))

	for i := 0; i < limit && len(allWords) > 0; i++ {
		var maxWord string
		var maxWordInfo Word
		for word, wordInfo := range allWords {
			if maxWord == "" || wordInfo.Freq > maxWordInfo.Freq {
				maxWord = word
				maxWordInfo = wordInfo
			}
		}

		words = append(words, maxWord)
		delete(allWords, maxWord)
	}

	return words
}

const allowedChangesQuotient = 3

func (e *Engine) CorrectWord(rawWord string, layout KeyboardLayoutNearbyKeys, limit int) []string {
	allowedChanges := len([]rune(rawWord)) / allowedChangesQuotient

	nearbyWords := e.GetNearbyWords(rawWord, allowedChanges, layout)
	words := make([]string, 0, len(nearbyWords))

	for i := 0; i < limit && len(nearbyWords) > 0; i++ {
		var minWord string
		var minWordInfo NearbyWordInfo
		for word, wordInfo := range nearbyWords {
			if minWord == "" ||
				wordInfo.Changes < minWordInfo.Changes ||
				(wordInfo.Changes == minWordInfo.Changes && wordInfo.Freq > minWordInfo.Freq) {
				minWord = word
				minWordInfo = wordInfo
			}
		}

		words = append(words, minWord)
		delete(nearbyWords, minWord)
	}

	return words
}

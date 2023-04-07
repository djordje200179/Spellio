package spellio

func (e *Engine) PredictWord(prefix string, count int) []string {
	allWords := e.GetWordsByPrefix(prefix)
	words := make([]string, 0, len(allWords))

	for i := 0; i < count && len(allWords) > 0; i++ {
		var maxWord string
		var maxOccurrence uint
		for word, occurrences := range allWords {
			if occurrences > maxOccurrence {
				maxWord = word
				maxOccurrence = occurrences
			}
		}

		words = append(words, maxWord)
		delete(allWords, maxWord)
	}

	return words
}

const allowedChangesQuotient = 3

func (e *Engine) CorrectWord(rawWord string, layout KeyboardLayoutNearbyKeys, limit uint) []string {
	allowedChanges := (len(rawWord) + allowedChangesQuotient - 1) / allowedChangesQuotient

	nearbyWords := e.GetNearbyWords(rawWord, uint(allowedChanges), layout)
	words := make([]string, 0, len(nearbyWords))

	for i := uint(0); i < limit && len(nearbyWords) > 0; i++ {
		var minWord string
		var minChanges uint
		for word, changes := range nearbyWords {
			if changes < minChanges || minWord == "" {
				minWord = word
				minChanges = changes
			}
		}

		words = append(words, minWord)
		delete(nearbyWords, minWord)
	}

	return words
}

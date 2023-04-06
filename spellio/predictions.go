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

package spellio

import (
	"golang.org/x/exp/slices"
)

func (e *Engine) CompleteWord(prefix string, limit int) []Word {
	allWords := e.GetWordsByPrefix(prefix)
	slices.SortFunc[Word](allWords, func(first, second Word) bool {
		return first.Freq > second.Freq
	})

	if len(allWords) > limit {
		return allWords[:limit]
	} else {
		return allWords
	}
}

const allowedChangesQuotient = 3

func (e *Engine) CorrectWord(rawWord string, layout KeyboardLayoutNearbyKeys, limit int) []Word {
	allowedChanges := len([]rune(rawWord)) / allowedChangesQuotient

	nearbyWords := e.GetNearbyWords(rawWord, allowedChanges, layout)
	slices.SortFunc[NearbyWordInfo](nearbyWords, func(first, second NearbyWordInfo) bool {
		if first.Changes < second.Changes {
			return true
		}

		if first.Changes == second.Changes {
			return first.Freq > second.Freq
		}

		return false
	})

	if limit > len(nearbyWords) {
		limit = len(nearbyWords)
	}

	words := make([]Word, limit)
	for i := 0; i < limit; i++ {
		words[i] = nearbyWords[i].Word
	}

	return words
}

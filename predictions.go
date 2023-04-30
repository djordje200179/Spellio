package spellio

import (
	"golang.org/x/exp/slices"
)

// CompleteWord returns a list of words that start with the given prefix.
//
// The list is sorted by frequency in descending order and
// can be limited by the given limit.
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

// CorrectWord returns a list of words that are similar to the given word.
//
// The list is sorted by the number of changes in ascending order and
// their frequency in descending order.
// The list can be limited by the given limit parameter.
//
// The number of allowed changes is calculated as the length of the word
// divided by the 3.
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

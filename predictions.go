package spellio

import (
	"slices"
)

// CompleteWord returns a list of words that start with the given prefix.
//
// The list is sorted by frequency in descending order and
// can be limited by the given limit.
func (e *Engine) CompleteWord(prefix string, limit int) []*Word {
	words := e.GetWordsByPrefix(prefix)
	slices.SortFunc(words, func(first, second *Word) int {
		return int(first.Freq) - int(second.Freq)
	})

	if len(words) > limit {
		return words[:limit]
	} else {
		return words
	}
}

const allowedChangesQuotient = 3

// CorrectWord returns a list of words that are similar to the given word.
//
// The list is sorted by the number of changes in ascending order
// and their frequency in descending order.
// The list can be limited by the given limit parameter.
//
// The number of allowed changes is calculated as the length
// of the word divided by the 3.
func (e *Engine) CorrectWord(rawWord string, layout KeyboardLayout, limit int) []*Word {
	maxChanges := len([]rune(rawWord)) / allowedChangesQuotient

	nearbyWords := e.GetNearbyWords(rawWord, maxChanges, layout)
	slices.SortFunc(nearbyWords, func(first, second NearbyWord) int {
		changesDiff := first.Changes - second.Changes
		if changesDiff != 0 {
			return changesDiff
		}

		return int(first.Freq) - int(second.Freq)
	})

	if limit > len(nearbyWords) {
		limit = len(nearbyWords)
	}

	words := make([]*Word, limit)
	for i := range limit {
		words[i] = nearbyWords[i].Word
	}

	return words
}

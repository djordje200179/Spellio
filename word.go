package spellio

// A Word is a word in the dictionary.
type Word struct {
	Freq uint // Frequency of the word in the dictionary.

	lastLetter *letterNode
}

// Empty returns true if the word is empty.
func (word Word) Empty() bool {
	return word.lastLetter == nil
}

// String returns the word as a string.
func (word Word) String() string {
	return word.lastLetter.getWord(nil)
}

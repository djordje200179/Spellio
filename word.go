package spellio

// A Word is a word in the dictionary.
type Word struct {
	Freq uint // Frequency of the word in the dictionary.

	lastLetter *letterNode
}

// String returns the word as a string.
func (w *Word) String() string {
	return w.lastLetter.getWord(nil)
}

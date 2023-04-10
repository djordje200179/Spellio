package spellio

type Word struct {
	Freq uint

	lastLetter *letterNode
}

func (word Word) Empty() bool {
	return word.lastLetter == nil
}

func (word Word) String() string {
	return word.lastLetter.getWord(nil)
}

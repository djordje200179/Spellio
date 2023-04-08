package spellio

import "fmt"

type Word struct {
	Freq uint
}

func (word Word) String() string {
	return fmt.Sprintf("%d", word.Freq)
}

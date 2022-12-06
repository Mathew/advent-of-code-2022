package signal

import (
	"strings"

	"github.com/mathew/advent-of-code-2022/internal/structures"
)

type Signal struct {
	sequence string
}

func NewSignal(sequence string) Signal {
	return Signal{sequence}
}

func (s Signal) StartMarkerLocation(length int) int {
	i, j := 0, length

	for true {
		if len(structures.Set(strings.Split(s.sequence[i:j], "")...)) == length {
			return j
		}

		i += 1
		j += 1
	}

	return 0
}

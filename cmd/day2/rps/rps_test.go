package rps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWinnerIs(t *testing.T) {
	assert.Equal(t, 6, scoreRPS(Rock, Scissor))
	assert.Equal(t, 6, scoreRPS(Scissor, Paper))
	assert.Equal(t, 0, scoreRPS(Rock, Paper))
	assert.Equal(t, 3, scoreRPS(Paper, Paper))
}

func TestStrategyGuideTotalScore(t *testing.T) {
	sg := NewStrategyGuide([][]string{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}, map[string]string{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
		"X": Rock,
		"Y": Paper,
		"Z": Scissor,
	}, map[string]int{
		Rock:    1,
		Paper:   2,
		Scissor: 3,
	})

	assert.Equal(t, 15, sg.TotalScore())
}

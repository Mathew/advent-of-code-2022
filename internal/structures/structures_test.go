package structures

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapInvert(t *testing.T) {

	d := map[string]int{
		"a": 1,
		"b": 2,
	}

	exp := map[int]string{
		1: "a",
		2: "b",
	}

	assert.Equal(t, exp, MapInvert(d))
}

func TestIntersection(t *testing.T) {
	assert.Equal(
		t,
		[]string{"b"},
		Intersection([]string{"a", "b", "c"}, []string{"A", "b", "C"}),
	)
}

func TestSequence(t *testing.T) {
	assert.Equal(t, []int{2, 3, 4}, Sequence(2, 4))
}

func TestSet(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, Set("a", "a", "b", "c", "a", "c"))
	assert.Equal(t, []string{"a", "b", "c"}, Set(strings.Split("aaabbbaaacccc", "")...))
}

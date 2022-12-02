package structures

import (
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

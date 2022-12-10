package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert.Equal(t, []string{"a"}, Filter(func(s string) bool { return s == "a" }, "a", "b", "c"))
}

func TestAny(t *testing.T) {
	result := Any(func(item bool) bool { return item == false }, true, true, true)
	assert.Equal(t, false, result)

	result = Any(func(item bool) bool { return item == false }, true, false, true)
	assert.Equal(t, true, result)
}

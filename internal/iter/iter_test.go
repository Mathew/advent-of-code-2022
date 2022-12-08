package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert.Equal(t, []string{"a"}, Filter(func(s string) bool { return s == "a" }, "a", "b", "c"))
}

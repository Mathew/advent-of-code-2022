package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignal(t *testing.T) {
	assert.Equal(t, 4, NewSignal("abcdefghijkklmno").StartMarkerLocation(4))
	assert.Equal(t, 5, NewSignal("bvwbjplbgvbhsrlpgdmjqwftvncz").StartMarkerLocation(4))
	assert.Equal(t, 6, NewSignal("nppdvjthqldpwncqszvftbrmjlhg").StartMarkerLocation(4))
	assert.Equal(t, 10, NewSignal("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg").StartMarkerLocation(4))
	assert.Equal(t, 11, NewSignal("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw").StartMarkerLocation(4))

	assert.Equal(t, 19, NewSignal("mjqjpqmgbljsphdztnvjfqwrcgsmlb").StartMarkerLocation(14))
}

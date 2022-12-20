package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPU(t *testing.T) {
	instructions := []Command{
		NewNoop(),
		NewAdd(3),
		NewAdd(-5),
	}

	cpu := NewCPU(instructions)

	assert.Equal(t, 1, cpu())
	assert.Equal(t, 1, cpu())
	assert.Equal(t, 1, cpu())
	assert.Equal(t, 4, cpu())
	assert.Equal(t, 4, cpu())
	assert.Equal(t, -1, cpu())
}

func TestCPUCounter(t *testing.T) {
	instructions := []Command{
		NewAdd(15), NewAdd(-11), NewAdd(6), NewAdd(-3), NewAdd(5), NewAdd(-1), NewAdd(-8), NewAdd(13), NewAdd(4), NewNoop(), NewAdd(-1), NewAdd(5), NewAdd(-1), NewAdd(5), NewAdd(-1), NewAdd(5), NewAdd(-1), NewAdd(5), NewAdd(-1), NewAdd(-35), NewAdd(1), NewAdd(24), NewAdd(-19), NewAdd(1), NewAdd(16), NewAdd(-11), NewNoop(), NewNoop(), NewAdd(21), NewAdd(-15), NewNoop(), NewNoop(), NewAdd(-3), NewAdd(9), NewAdd(1), NewAdd(-3), NewAdd(8), NewAdd(1), NewAdd(5), NewNoop(), NewNoop(), NewNoop(), NewNoop(), NewNoop(), NewAdd(-36), NewNoop(), NewAdd(1), NewAdd(7), NewNoop(), NewNoop(), NewNoop(), NewAdd(2), NewAdd(6), NewNoop(), NewNoop(), NewNoop(), NewNoop(), NewNoop(), NewAdd(1), NewNoop(), NewNoop(), NewAdd(7), NewAdd(1), NewNoop(), NewAdd(-13), NewAdd(13), NewAdd(7), NewNoop(), NewAdd(1), NewAdd(-33), NewNoop(), NewNoop(), NewNoop(), NewAdd(2), NewNoop(), NewNoop(), NewNoop(), NewAdd(8), NewNoop(), NewAdd(-1), NewAdd(2), NewAdd(1), NewNoop(), NewAdd(17), NewAdd(-9), NewAdd(1), NewAdd(1), NewAdd(-3), NewAdd(11), NewNoop(), NewNoop(), NewAdd(1), NewNoop(), NewAdd(1), NewNoop(), NewNoop(), NewAdd(-13), NewAdd(-19), NewAdd(1), NewAdd(3), NewAdd(26), NewAdd(-30), NewAdd(12), NewAdd(-1), NewAdd(3), NewAdd(1), NewNoop(), NewNoop(), NewNoop(), NewAdd(-9), NewAdd(18), NewAdd(1), NewAdd(2), NewNoop(), NewNoop(), NewAdd(9), NewNoop(), NewNoop(), NewNoop(), NewAdd(-1), NewAdd(2), NewAdd(-37), NewAdd(1), NewAdd(3), NewNoop(), NewAdd(15), NewAdd(-21), NewAdd(22), NewAdd(-6), NewAdd(1), NewNoop(), NewAdd(2), NewAdd(1), NewNoop(), NewAdd(-10), NewNoop(), NewNoop(), NewAdd(20), NewAdd(1), NewAdd(2), NewAdd(2), NewAdd(-6), NewAdd(-11), NewNoop(), NewNoop(), NewNoop()}

	cpu := CycleCounter([]int{20, 60, 100, 140, 180, 220}, NewCPU(instructions))
	assert.Equal(t, 420, cpu())
	assert.Equal(t, 1140, cpu())
	assert.Equal(t, 1800, cpu())
	assert.Equal(t, 2940, cpu())
	assert.Equal(t, 2880, cpu())
	assert.Equal(t, 3960, cpu())
}

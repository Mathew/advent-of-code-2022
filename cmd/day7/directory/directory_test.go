package directory

import (
	"testing"

	"github.com/mathew/advent-of-code-2022/internal/iter"
	"github.com/mathew/advent-of-code-2022/internal/maths"
	"github.com/mathew/advent-of-code-2022/internal/structures"
	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	terminalOutput := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k ",
	}

	assert.Equal(t,
		map[string]int{"/": 48381165, "/.a": 94853, "/.d": 24933642, "/.a.e": 584},
		Evaluate(terminalOutput))

	r := maths.Sum(
		iter.Filter(
			func(i int) bool { return i < 100001 },
			structures.MapValues(Evaluate(terminalOutput))...,
		)...,
	)

	assert.Equal(t, 95437, r)
}

func TestFolderNamingCollision(t *testing.T) {
	terminalOutput := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"dir a",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k ",
		"$ cd a",
		"$ ls",
		"1 a.dupe",
	}

	assert.Equal(t,
		map[string]int{
			"/":     48381166,
			"/.a":   94853,
			"/.d":   24933643,
			"/.a.e": 584,
			"/.d.a": 1,
		},
		Evaluate(terminalOutput))

	r := maths.Sum(
		iter.Filter(
			func(i int) bool { return i < 100001 },
			structures.MapValues(Evaluate(terminalOutput))...,
		)...,
	)

	assert.Equal(t, 95438, r)
}

func TestSmallestViableDirectory(t *testing.T) {
	used := 48381165
	directories := map[string]int{"/": 48381165, "/.a": 94853, "/.d": 24933642, "/.a.e": 584}

	assert.Equal(t,
		24933642,
		SmallestViableDirectory(used, structures.MapValues(directories)))
}

package directory

import (
	"log"
	"strconv"
	"strings"
)

func Evaluate(output []string) map[string]int {
	directorySizes := map[string]int{}
	currentPath := []string{}

	for _, line := range output {
		parts := strings.Split(line, " ")

		// ignore directory listings.
		if parts[0] == "dir" {
			continue
		}

		if parts[0] == string("$") {
			// handle directory changes
			if parts[1] == "cd" {
				if parts[2] == ".." {
					currentPath = currentPath[:len(currentPath)-1]
				} else {
					currentPath = append(currentPath, parts[2])
				}
			}

			// just ignore list commands.
			continue
		}

		for i := 0; i < len(currentPath); i++ {
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalf("Unable to convert %s", parts[0])
			}

			directorySizes[strings.Join(currentPath[0:i+1], ".")] += size
		}
	}

	return directorySizes
}

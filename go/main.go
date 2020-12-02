package adventofcode2020

import (
	"fmt"
	"os"

	"github.com/jghiloni/adventofcode2020/day1"
	"github.com/jghiloni/adventofcode2020/day2"
)

// Main does the unit of work and can be tested
func Main(args []string) error {

	// This structure (app day part) could be accomplished with one of many flag
	// parsing libraries that exist for Go, but it would be overkill
	commands := map[string]DailyExercise{
		"day1": day1.Exercise{},
		"day2": day2.Exercise{},
	}

	if len(args) < 3 {
		return ErrUsage
	}

	command, ok := commands[args[1]]
	if !ok {
		return fmt.Errorf("Unrecognized command %q", args[1])
	}

	part := args[2]
	method := command.Part1

	if part == "part2" {
		method = command.Part2
	}

	return method(os.Stdin, os.Stdout)
}

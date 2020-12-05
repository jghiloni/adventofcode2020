package adventofcode2020

import (
	"fmt"
	"os"
	"path"

	"github.com/jghiloni/adventofcode2020/day1"
	"github.com/jghiloni/adventofcode2020/day2"
	"github.com/jghiloni/adventofcode2020/day3"
	"github.com/jghiloni/adventofcode2020/day4"
	"github.com/jghiloni/adventofcode2020/day5"
	"github.com/jghiloni/adventofcode2020/errors"
)

// Main does the unit of work and can be tested
func Main(args []string) error {

	// This structure (app day part) could be accomplished with one of many flag
	// parsing libraries that exist for Go, but it would be overkill
	commands := map[string]DailyExercise{
		"day1": day1.Exercise{},
		"day2": day2.Exercise{},
		"day3": day3.Exercise{},
		"day4": day4.Exercise{},
		"day5": day5.Exercise{},
	}

	if len(args) < 3 {
		return errors.ErrUsage
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

	input := os.Stdin
	wd, _ := os.Getwd()
	var err error
	if _, err = os.Stat(path.Join(wd, "..", args[1], "input")); err == nil {
		input, err = os.Open(path.Join(wd, "..", args[1], "input"))
		if err != nil {
			return err
		}
	}

	return method(input, os.Stdout)
}

package adventofcode2020

import (
	"fmt"
	"os"
	"path"

	"github.com/jghiloni/adventofcode2020/day1"
	"github.com/jghiloni/adventofcode2020/day10"
	"github.com/jghiloni/adventofcode2020/day11"
	"github.com/jghiloni/adventofcode2020/day12"
	"github.com/jghiloni/adventofcode2020/day13"
	"github.com/jghiloni/adventofcode2020/day14"
	"github.com/jghiloni/adventofcode2020/day15"
	"github.com/jghiloni/adventofcode2020/day16"
	"github.com/jghiloni/adventofcode2020/day17"
	"github.com/jghiloni/adventofcode2020/day18"
	"github.com/jghiloni/adventofcode2020/day19"
	"github.com/jghiloni/adventofcode2020/day2"
	"github.com/jghiloni/adventofcode2020/day20"
	"github.com/jghiloni/adventofcode2020/day21"
	"github.com/jghiloni/adventofcode2020/day22"
	"github.com/jghiloni/adventofcode2020/day23"
	"github.com/jghiloni/adventofcode2020/day24"
	"github.com/jghiloni/adventofcode2020/day25"
	"github.com/jghiloni/adventofcode2020/day3"
	"github.com/jghiloni/adventofcode2020/day4"
	"github.com/jghiloni/adventofcode2020/day5"
	"github.com/jghiloni/adventofcode2020/day6"
	"github.com/jghiloni/adventofcode2020/day7"
	"github.com/jghiloni/adventofcode2020/day8"
	"github.com/jghiloni/adventofcode2020/day9"
	"github.com/jghiloni/adventofcode2020/errors"
)

// Main does the unit of work and can be tested
func Main(args []string) error {

	// This structure (app day part) could be accomplished with one of many flag
	// parsing libraries that exist for Go, but it would be overkill
	commands := map[string]DailyExercise{
		"day1":  day1.Exercise{},
		"day2":  day2.Exercise{},
		"day3":  day3.Exercise{},
		"day4":  day4.Exercise{},
		"day5":  day5.Exercise{},
		"day6":  day6.Exercise{},
		"day7":  day7.Exercise{},
		"day8":  day8.Exercise{},
		"day9":  day9.Exercise{},
		"day10": day10.Exercise{},
		"day11": day11.Exercise{},
		"day12": day12.Exercise{},
		"day13": day13.Exercise{},
		"day14": day14.Exercise{},
		"day15": day15.Exercise{},
		"day16": day16.Exercise{},
		"day17": day17.Exercise{},
		"day18": day18.Exercise{},
		"day19": day19.Exercise{},
		"day20": day20.Exercise{},
		"day21": day21.Exercise{},
		"day22": day22.Exercise{},
		"day23": day23.Exercise{},
		"day24": day24.Exercise{},
		"day25": day25.Exercise{},
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

package adventofcode2020

import (
	"errors"
	"io"
)

// DailyExercise represents each day's exercise
type DailyExercise interface {
	Part1(input io.Reader, output io.Writer) error
	Part2(input io.Reader, output io.Writer) error
}

// ErrUsage will be thrown in Main if command line args are wrong
var ErrUsage = errors.New(`Usage: adventofcode2020 dayN [part1|part2]`)

// ErrUnimplemented will be returned if the method is unimplemented
var ErrUnimplemented = errors.New(`unimplemented`)

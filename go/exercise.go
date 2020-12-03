package adventofcode2020

import (
	"io"
)

// DailyExercise represents each day's exercise
type DailyExercise interface {
	Part1(input io.Reader, output io.Writer) error
	Part2(input io.Reader, output io.Writer) error
}

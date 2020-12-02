package day2

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// Exercise is the implementation of DailyExercise for 2020-12-02
type Exercise struct{}

// Part1 reads from the input and determines if the password, represented in the line
//     a-b c: password
// has between a and b instances of character c in the password, and displays the number
// of lines that meet this criteria
func (e Exercise) Part1(input io.Reader, output io.Writer) error {
	inputBytes, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}

	validPasswords := 0
	lines := strings.Split(string(inputBytes), "\n")
	var (
		password string
		c        string
		min      int
		max      int
	)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		fmt.Sscanf(line, "%d-%d %s %s", &min, &max, &c, &password)
		condition := rune(c[0])

		total := 0
		for _, char := range line {
			if char == condition {
				total++
			}
		}

		if total >= min && total <= max {
			validPasswords++
		}
	}

	fmt.Fprintln(output, validPasswords)
	return nil
}

// Part2 reads the same input and determines if the given password has the
// character c at position a XOR position b in the password, and returns the
// number of password that meet this criterion.
func (e Exercise) Part2(input io.Reader, output io.Writer) error {
	inputBytes, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}

	validPasswords := 0
	lines := strings.Split(string(inputBytes), "\n")
	var (
		password string
		c        string
		first    int64
		second   int64
	)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		fmt.Sscanf(line, "%d-%d %s %s", &first, &second, &c, &password)
		condition := c[0]

		var (
			match1 int
			match2 int
		)

		if password[first-1] == condition {
			match1 = 1
		}

		if password[second-1] == condition {
			match2 = 1
		}

		if (match1 ^ match2) == 1 {
			validPasswords++
		}
	}

	fmt.Fprintln(output, validPasswords)
	return nil
}

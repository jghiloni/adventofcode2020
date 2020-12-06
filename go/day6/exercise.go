package day6

import (
	"bufio"
	"fmt"
	"io"
)

// Exercise for 2020-12-06
type Exercise struct{}

func (e Exercise) Part1(in io.Reader, out io.Writer) error {
	groups, err := parseInput(in)
	if err != nil {
		return err
	}

	total := 0
	for _, group := range groups {
		total += group.countAllYesAnswers()
	}

	fmt.Fprintln(out, total)
	return nil
}

func (e Exercise) Part2(in io.Reader, out io.Writer) error {
	groups, err := parseInput(in)
	if err != nil {
		return err
	}

	total := 0
	for _, group := range groups {
		total += group.countCommonYesAnswers()
	}

	fmt.Fprintln(out, total)
	return nil
}

type groupAnswers []string

func (g groupAnswers) countAllYesAnswers() int {
	yeses := map[rune]bool{}
	for _, line := range g {
		for _, char := range line {
			yeses[char] = true
		}
	}

	return len(yeses)
}

func (g groupAnswers) countCommonYesAnswers() int {
	yeses := map[rune]int{}
	for _, line := range g {
		for _, char := range line {
			if _, ok := yeses[char]; !ok {
				yeses[char] = 0
			}

			yeses[char]++
		}
	}

	total := 0
	groupSize := len(g)
	for _, count := range yeses {
		if count == groupSize {
			total++
		}
	}

	return total
}

func parseInput(in io.Reader) ([]groupAnswers, error) {
	scanner := bufio.NewScanner(in)

	groups := make([]groupAnswers, 0, 100)
	group := groupAnswers{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groups = append(groups, group)
			group = groupAnswers{}
			continue
		}

		group = append(group, line)
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}

	return groups, scanner.Err()
}

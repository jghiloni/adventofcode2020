package day5

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Exercise for 2020-12-05
type Exercise struct{}

func (e Exercise) Part1(in io.Reader, out io.Writer) error {
	lines, err := parseInput(in)
	if err != nil {
		return err
	}

	highestSeatPosition := 0
	for _, line := range lines {
		fmt.Fprintln(out, "Checking code", line)

		seatPosition := getSeat(line).getSeatID()
		if seatPosition > highestSeatPosition {
			highestSeatPosition = seatPosition
		}
	}

	fmt.Fprintln(out, highestSeatPosition)
	return nil
}

func (e Exercise) Part2(in io.Reader, out io.Writer) error {
	lines, err := parseInput(in)
	if err != nil {
		return err
	}

	mySeat := seat{}
	seats := make([]seat, len(lines))
	for i := range lines {
		seats[i] = getSeat(lines[i])
	}

	sort.Slice(seats, func(i, j int) bool {
		s1, s2 := seats[i], seats[j]

		if s1.row < s2.row {
			return true
		}

		if s1.row > s2.row {
			return false
		}

		return s1.col < s2.col
	})

	i := 0
	for i = range seats {
		if seats[i].row > 0 {
			break
		}
	}

	for ; i < len(seats)-8; i += 8 {
		left := seats[i]
		right := seats[i+7]

		if left.row != right.row {
			for j := i; j < i+7; j++ {
				if seats[j].col-seats[j+1].col == -2 {
					mySeat = seat{row: seats[j].row, col: seats[j].col + 1}
					break
				}
			}
		}
	}

	fmt.Fprintln(out, mySeat.getSeatID())

	return nil
}

func parseInput(in io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(in)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	sort.Slice(lines, func(i, j int) bool {
		s1, s2 := lines[i], lines[j]

		row1, row2 := s1[:7], s2[:7]
		switch strings.Compare(row1, row2) {
		case -1:
			return true
		case 1:
			return false
		default:
			return strings.Compare(s1[7:], s2[7:]) > 0
		}
	})

	return lines, nil
}

func getSeat(line string) seat {
	maxRow := 127
	maxCol := 7

	minRow := 0
	minCol := 0

	for i := 0; i < 7; i++ {
		char := line[i]
		if char == 'F' {
			maxRow = minRow + ((maxRow - minRow) / 2)
			continue
		}

		minRow = minRow + (((maxRow - minRow) + 1) / 2)
	}

	row := minRow
	if line[6] == 'B' {
		row = maxRow
	}

	for i := 7; i < 10; i++ {
		char := line[i]
		if char == 'L' {
			maxCol = minCol + ((maxCol - minCol) / 2)
			continue
		}

		minCol = minCol + (((maxCol - minCol) + 1) / 2)
	}

	col := minCol
	if line[9] == 'R' {
		col = maxCol
	}

	return seat{row: row, col: col}
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}

	return false
}

type seat struct {
	row int
	col int
}

func (s seat) getSeatID() int {
	return (s.row * 8) + s.col
}

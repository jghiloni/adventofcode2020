package day3

import (
	"fmt"
	"io"
	"io/ioutil"
)

// Exercise is the exercise for 2020-12-04
type Exercise struct{}

// Part1 says Given the input grid, determine how many trees (marked by #)
// will be traveled through on the way to the bottom of the grid if eavh move
// is x + 3, y + 1, given the top left is (0,0)
func (e Exercise) Part1(input io.Reader, output io.Writer) error {
	grid, err := readInputGrid(input)
	if err != nil {
		return err
	}

	treeCount := countTrees(grid, 3, 1)

	fmt.Fprintln(output, treeCount)
	return nil
}

// Part2 requires using the algorithm from Part1 to determine the tree count
// given several different paths, and returning the product of all those counts
func (e Exercise) Part2(input io.Reader, output io.Writer) error {
	grid, err := readInputGrid(input)
	if err != nil {
		return err
	}

	counts := []int{}
	counts = append(counts, countTrees(grid, 1, 1))
	counts = append(counts, countTrees(grid, 3, 1))
	counts = append(counts, countTrees(grid, 5, 1))
	counts = append(counts, countTrees(grid, 7, 1))
	counts = append(counts, countTrees(grid, 1, 2))

	product := 1
	for _, count := range counts {
		product *= count
	}

	fmt.Fprintln(output, product)
	return nil
}

func readInputGrid(in io.Reader) ([][]bool, error) {
	returnVal := make([][]bool, 0, 350)

	body, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	yIndex := 0
	lineLen := 0

	for i, b := range body {
		if yIndex == len(returnVal) {
			returnVal = append(returnVal, make([]bool, 0, 32))
		}

		if b == '\n' {
			if lineLen == 0 {
				lineLen = i
			}

			yIndex++
			continue
		}

		returnVal[yIndex] = append(returnVal[yIndex], (b == '#'))
	}

	return returnVal, nil
}

func countTrees(grid [][]bool, xStep int, yStep int) int {
	yIndex := 0
	xIndex := 0
	xModulus := len(grid[0])
	treeCount := 0
	for yIndex < len(grid)-yStep {
		if grid[yIndex+yStep][(xIndex+xStep)%xModulus] {
			treeCount++
		}

		yIndex += yStep
		xIndex += xStep
	}

	return treeCount
}

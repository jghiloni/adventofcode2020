package day9

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	goerr "errors"
)

type Exercise struct{}

func (e Exercise) Part1(in io.Reader, out io.Writer) error {
	nums, err := parseInput(in)
	if err != nil {
		return err
	}

	preambleLength := 25
	for i := preambleLength; i < len(nums); i++ {
		prevNums := nums[i-preambleLength : i]
		if !containsSumParts(prevNums, nums[i]) {
			fmt.Fprintln(out, nums[i])
			return nil
		}
	}

	return goerr.New("can't find an invalid number")
}

func (e Exercise) Part2(in io.Reader, out io.Writer) error {
	b := &bytes.Buffer{}
	if _, err := io.Copy(b, in); err != nil {
		return err
	}
	inputBytes := b.Bytes()

	p1out := &bytes.Buffer{}
	if err := e.Part1(bytes.NewBuffer(inputBytes), p1out); err != nil {
		return err
	}

	nums, err := parseInput(bytes.NewBuffer(inputBytes))
	if err != nil {
		return err
	}

	corruptedNumber, err := strconv.Atoi(strings.TrimSpace(p1out.String()))
	if err != nil {
		return err
	}

	total := findWeakness(nums, corruptedNumber)
	fmt.Fprintln(out, total)

	return nil
}

func parseInput(in io.Reader) ([]int, error) {
	numbers := []int{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, n)
	}

	return numbers, scanner.Err()
}

func containsSumParts(nums []int, total int) bool {
	for _, a1 := range nums {
		for _, a2 := range nums {
			if a1 == a2 {
				continue
			}

			if a1+a2 == total {
				return true
			}
		}
	}

	return false
}

func findWeakness(nums []int, corruptedNumber int) int {
	startIdx := 0
	endIdx := 0
	for i := range nums {
		startIdx = i
		j := i + 1
		total := nums[i]
		for j < len(nums) {
			if total+nums[j] == corruptedNumber {
				endIdx = j
				break
			}

			total += nums[j]
			j++
		}

		if endIdx != 0 {
			break
		}
	}

	min := nums[startIdx]
	max := 0
	for _, n := range nums[startIdx : endIdx+1] {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return min + max
}

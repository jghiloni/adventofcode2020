package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)

// Exercise is the implementation of DailyExercise for 2020-12-01
type Exercise struct{}

// Part1 requires us to find the two numbers in the expense report that add up
// to 2020, and return their product
func (e Exercise) Part1(input io.Reader, output io.Writer) error {
	numbers, err := getNumbers(input)
	if err != nil {
		return err
	}

	var (
		num1 int
		num2 int
	)

MainLoop:
	for i := range numbers {
		for j := range numbers {
			if i == j {
				continue
			}

			if numbers[i]+numbers[j] == 2020 {
				num1 = numbers[i]
				num2 = numbers[j]

				break MainLoop
			}

		}

		if num1 != 0 && num2 != 0 {
			break
		}
	}

	if num1 == 0 && num2 == 0 {
		log.Fatal("couldn't find numbers")
	}

	log.Printf("the numbers are %d and %d\n", num1, num2)
	fmt.Fprintln(output, num1*num2)

	return nil
}

// Part2 is like Part1, but for three items, not two
func (e Exercise) Part2(input io.Reader, output io.Writer) error {
	numbers, err := getNumbers(input)
	if err != nil {
		return err
	}

	var (
		num1 int
		num2 int
		num3 int
	)

MainLoop:
	for i := range numbers {
		for j := range numbers {
			if i == j {
				continue
			}
			for k := range numbers {
				if i == k || j == k {
					continue
				}

				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					num1 = numbers[i]
					num2 = numbers[j]
					num3 = numbers[k]

					break MainLoop
				}
			}
		}

		if num1 != 0 && num2 != 0 && num3 != 0 {
			break
		}
	}

	if num1 == 0 && num2 == 0 && num3 == 0 {
		log.Fatal("couldn't find numbers")
	}

	log.Printf("the numbers are %d, %d, and %d\n", num1, num2, num3)
	fmt.Fprintln(output, num1*num2*num3)

	return nil
}

func getNumbers(input io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(input)

	numbers := make([]int, 0, 500)
	for scanner.Scan() {
		str := scanner.Text()
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

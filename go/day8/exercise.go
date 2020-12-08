package day8

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
)

type Exercise struct{}

type opcode struct {
	name     string
	value    int
	executed bool
}

var errCycle = errors.New("instruction cycle detected")

func (e Exercise) Part1(in io.Reader, out io.Writer) error {
	opcodes, err := parseInput(in)
	if err != nil {
		return err
	}

	accumulator := 0
	accumulator, err = findOpcodeCycle(opcodes, accumulator, 0)
	if errors.Is(err, errCycle) {
		fmt.Fprintln(out, accumulator)
	}
	return err
}

func (e Exercise) Part2(in io.Reader, out io.Writer) error {
	codes, err := parseInput(in)
	if err != nil {
		return err
	}

	for line := range codes {
		if codes[line].name == "jmp" && codes[line].value < 0 {
			log.Printf("changing line %d from jmp %d to nop", line, codes[line].value)
			codes[line].name = "nop"

			resetExecution(codes)
			accumulator, err := findOpcodeCycle(codes, 0, 0)
			if errors.Is(err, errCycle) {
				codes[line].name = "jmp"
				continue
			}

			if err == nil {
				fmt.Fprintln(out, accumulator)
				return nil
			}

			return err
		}
	}

	return errors.New("could not find a version of the code that doesn't have a cycle")
}

func parseInput(in io.Reader) ([]opcode, error) {
	scanner := bufio.NewScanner(in)
	codes := []opcode{}
	for scanner.Scan() {
		o := opcode{}
		fmt.Sscanf(scanner.Text(), "%s %d", &o.name, &o.value)

		codes = append(codes, o)
	}

	return codes, scanner.Err()
}

func findOpcodeCycle(codes []opcode, accumulator int, line int) (int, error) {
	code := codes[line]
	if code.executed {
		return accumulator, errCycle
	}

	codes[line].executed = true

	switch code.name {
	case "acc":
		line++
		accumulator += code.value
	case "jmp":
		line += code.value
	case "nop":
		line++
	}

	if line >= len(codes) {
		return accumulator, nil
	}

	return findOpcodeCycle(codes, accumulator, line)
}

func resetExecution(codes []opcode) {
	for i := range codes {
		codes[i].executed = false
	}
}

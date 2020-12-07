package day7

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Exercise struct{}

func (e Exercise) Part1(in io.Reader, out io.Writer) error {
	r, err := newRuleSet(in)
	if err != nil {
		return err
	}

	count := len(r.canContain("shiny gold"))
	fmt.Fprintln(out, count)

	return nil
}

func (e Exercise) Part2(in io.Reader, out io.Writer) error {
	r, err := newRuleSet(in)
	if err != nil {
		return err
	}

	total, err := r.countContents("shiny gold")
	if err != nil {
		return err
	}
	fmt.Fprint(out, total)

	return nil
}

type ruleSet map[string][]string

func newRuleSet(in io.Reader) (ruleSet, error) {
	rs := ruleSet{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), ".")

		// normalize line
		line = strings.ReplaceAll(line, "bags", "bag")
		line = strings.ReplaceAll(line, "bag", "")
		line = strings.TrimSpace(line)

		rule := strings.Split(line, " contain ")
		if rule[1] == "no other" {
			continue
		}

		contents := strings.Split(rule[1], ", ")
		rs[strings.TrimSpace(rule[0])] = contents
	}

	return rs, scanner.Err()
}

func (r ruleSet) canContain(bagType string) []string {
	containers := map[string]bool{}
	for bag, contents := range r {
		if r.contains(contents, bagType) {
			containers[bag] = true
		}
	}

	bags := make([]string, 0, len(containers))
	for bag := range containers {
		bags = append(bags, bag)
	}

	return bags
}

func (r ruleSet) contains(contents []string, bagType string) bool {
	if len(contents) == 0 {
		return false
	}

	for _, subBag := range contents {
		numType := strings.SplitN(subBag, " ", 2)

		if len(numType) == 1 {
			return false
		}

		c := strings.TrimSpace(numType[1])

		if c == bagType {
			return true
		}

		cbag, ok := r[c]
		if !ok {
			continue
		}

		if r.contains(cbag, bagType) {
			return true
		}
	}

	return false
}

func (r ruleSet) countContents(bagType string) (int64, error) {
	bagContents, ok := r[bagType]
	if !ok {
		return 0, nil
	}

	total := int64(0)
	for _, innerBag := range bagContents {
		numType := strings.SplitN(innerBag, " ", 2)
		count, err := strconv.ParseInt(strings.TrimSpace(numType[0]), 10, 0)
		if err != nil {
			log.Println(err)
			return 0, err
		}

		subBagType := strings.TrimSpace(numType[1])
		subCount, err := r.countContents(subBagType)
		if err != nil {
			return 0, err
		}
		total += (count * subCount) + count
	}

	return total, nil
}

package day4

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// Exercise for 2020-12-04
type Exercise struct{}

// Part1 asks the developer to read the input and determine
// the number of valid passports. Valid passports have fields
// in the input in key:value format and are separated by blank
// lines. The fields:
// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID) *OPTIONAL*
func (e Exercise) Part1(in io.Reader, out io.Writer) error {
	passports, err := getPassports(in)
	if err != nil {
		return err
	}

	total := 0
	for _, p := range passports {
		if p.validate("cid") {
			total++
		}
	}

	fmt.Fprintln(out, total)
	return nil
}

// Part2 asks us to further validate the passport values. See
// the passport.ValidateFields method for the validation rules.
func (e Exercise) Part2(in io.Reader, out io.Writer) error {
	passports, err := getPassports(in)
	if err != nil {
		return err
	}

	total := 0
	for _, p := range passports {
		if p.validateFields() {
			total++
		}
	}

	fmt.Fprintln(out, total)
	return nil
}

func getPassports(in io.Reader) ([]passport, error) {
	scanner := bufio.NewScanner(in)

	passports := make([]passport, 0, 100)
	currentPassport := passport(map[string]string{})
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			if len(currentPassport) > 0 {
				passports = append(passports, currentPassport)
			}
			currentPassport = passport(map[string]string{})
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			kv := strings.Split(field, ":")
			if len(kv) == 1 {
				return nil, fmt.Errorf("invalid field found %q", field)
			}

			currentPassport[kv[0]] = kv[1]
		}
	}

	if len(currentPassport) > 0 {
		passports = append(passports, currentPassport)
	}

	return passports, scanner.Err()
}

type passport map[string]string

var passportFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

func (p passport) validate(ignoredFields ...string) bool {
	// fill ignored fields with an irrelevant value if not already set
	for _, ignored := range ignoredFields {
		if _, ok := p[ignored]; !ok {
			p[ignored] = "SHIM"
		}
	}

	return len(p) == len(passportFields)
}

/* validateFields validates the following rules:
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
*/
func (p passport) validateFields() bool {
	if !p.validate("cid") {
		return false
	}

	eyecolors := " amb blu brn gry grn hzl oth "

	total := 1
	for key, val := range p {
		switch key {
		case "cid":
			continue
		case "byr":
			if validateNumber(val, 4, 1920, 2002) {
				total++
			}
		case "iyr":
			if validateNumber(val, 4, 2010, 2020) {
				total++
			}
		case "eyr":
			if validateNumber(val, 4, 2020, 2030) {
				total++
			}
		case "pid":
			if validateNumber(val, 9, 0, 0) {
				total++
			}
		case "ecl":
			if strings.Contains(eyecolors, " "+val+" ") {
				total++
			}
		case "hgt":
			if len(val) < 3 {
				continue
			}

			num := val[:len(val)-2]
			unit := val[len(val)-2:]

			valid := false
			if unit == "in" {
				valid = validateNumber(num, 2, 59, 76)
			} else if unit == "cm" {
				valid = validateNumber(num, 3, 150, 193)
			}

			if valid {
				total++
			}
		case "hcl":
			var hex string
			fmt.Sscanf(val, "#%s", &hex)

			if len(hex) != 6 {
				continue
			}

			num, err := strconv.ParseInt(hex, 16, 0)
			if err != nil {
				continue
			}

			if num >= 0 && num <= 0xFFFFFF {
				total++
			}
		}
	}

	return total == len(p)
}

func validateNumber(num string, digits int, min int64, max int64) bool {
	if max == 0 {
		max = int64(math.Pow10(digits)) - 1
	}

	if len(num) != digits {
		return false
	}

	n, err := strconv.ParseInt(num, 10, 0)
	if err != nil {
		return false
	}

	return n >= min && n <= max
}

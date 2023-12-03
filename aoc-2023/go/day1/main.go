package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	data, err := os.ReadFile("input")
	if err != nil {
		panic("can't read the file")
	}

	replaceAllNumberStrings := regexp.MustCompile(getNumberPattern(numbers))
	numbersOnlyExp := regexp.MustCompile("[^0-9]+")

	sum := 0
	for _, line := range strings.Split(string(data), "\n") {
		replaced := replaceAllNumberStrings.ReplaceAllStringFunc(line, func(match string) string {
			return numbers[match]
		})
		result := numbersOnlyExp.ReplaceAllString(replaced, "")
		digits := strings.Split(result, "")
		num, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
		if err != nil {
			panic("bad bad")
		}
		sum += num
	}

	println(sum)
}

func getNumberPattern(wordToDigit map[string]string) string {
	pattern := ""
	for word := range wordToDigit {
		if pattern != "" {
			pattern += "|"
		}
		pattern += word
	}
	return pattern
}

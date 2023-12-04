package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func readInputToLines(fileName string) []string {
	data, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal("Couldn't read the input")
	}

	var result []string
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		result = append(result, strings.Trim(line, " "))
	}
	return result
}

func parseCard(line string) ([]string, []string) {
	idToResults := strings.Split(line, ":")

	results := strings.Split(idToResults[1], "|")
	winningNumbers, ourNumbers := results[0], results[1]

	winningNumbers = strings.Trim(winningNumbers, " ")
	ourNumbers = strings.Trim(ourNumbers, " ")
	winningNumbersList := filter(strings.Split(winningNumbers, " "), func(s string) bool {
		return strings.Trim(s, " ") != ""
	})
	ourNumbersList := filter(strings.Split(ourNumbers, " "), func(s string) bool {
		return strings.Trim(s, " ") != ""
	})

	return winningNumbersList, ourNumbersList
}

func calculateScore(matches int) int {
	score := 0
	if matches == 1 {
		score += 1
	} else if matches == 2 {
		score += 2
	} else if matches > 2 {
		score = 1
		for i := 1; i < matches; i++ {
			score = score * 2
		}
	}
	return score
}

func findMatchingNumbers(winners []string, ours []string) []string {
	var matches []string
	for _, winningNumber := range winners {
		if slices.Contains(ours, winningNumber) {
			matches = append(matches, winningNumber)
		}
	}
	return matches
}

func addNextMatch(pos int, times int, matches [][]string, total *int) {
	if times == 0 {
		return
	}

	for i := 1; i <= times; i++ {
		*total += 1
		nextMatch := matches[pos+i]
		addNextMatch(pos+i, len(nextMatch), matches, total)
	}
}

func main() {
	lines := readInputToLines("input")

	totalScore := 0
	for _, line := range lines {
		winningNumbers, ourNumbers := parseCard(line)

		matches := findMatchingNumbers(winningNumbers, ourNumbers)

		cardScore := calculateScore(len(matches))

		totalScore += cardScore
	}
	// fmt.Println("ending amount ", cards)
	fmt.Println("part 1: ", totalScore)

	var allMatches [][]string
	for _, line := range lines {
		winningNumbers, ourNumbers := parseCard(line)

		allMatches = append(allMatches, findMatchingNumbers(winningNumbers, ourNumbers))
	}

	total := len(allMatches)
	for idx, matchesForCard := range allMatches {
		addNextMatch(idx, len(matchesForCard), allMatches, &total)
	}

	fmt.Println("part 2: ", total)
}

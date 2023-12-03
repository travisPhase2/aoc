package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("coudln't read the file")
	}
	lines := strings.Split(string(input), "\n")

	var games []Game
	for _, line := range lines {
		separated := strings.Split(line, ":")
		id := strings.Split(separated[0], " ")[1]
		game := Game{
			Id: id,
		}
		setStrings := strings.Split(separated[1], ";")
		game.AddSets(setStrings)
		games = append(games, game)
	}

	search := Set{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	var foundGames []string
	for _, game := range games {
		ok := game.IsPotentialGameForSet(&search)
		if ok {
			foundGames = append(foundGames, game.Id)
		}
	}

	idSum := 0
	for _, game := range foundGames {
		conv, err := strconv.Atoi(game)
		if err != nil {
			panic("uh oh")
		}

		idSum += conv
	}

	fmt.Println("Part 1: ", idSum)

	powerSum := 0
	for _, game := range games {
		s := game.LowestPossibleDice()
		powerSum += (s.Red * s.Green * s.Blue)
	}

	fmt.Println("Part 2: ", powerSum)
}

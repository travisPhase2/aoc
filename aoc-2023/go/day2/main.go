package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id   string
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (g *Game) AddSets(setLines []string) {
	for _, setStr := range setLines {
		set := Set{
			Red:   0,
			Green: 0,
			Blue:  0,
		}
		rolls := strings.Split(setStr, ",")
		for _, roll := range rolls {
			amountToColor := strings.Split(strings.Trim(roll, " "), " ")
			color := amountToColor[1]
			amount, err := strconv.Atoi(amountToColor[0])
			if err != nil {
				panic("coulnd't parase amount for round")
			}

			switch color {
			case "red":
				set.Red += amount
			case "blue":
				set.Blue += amount
			case "green":
				set.Green += amount
			}
		}
		g.Sets = append(g.Sets, set)
	}
}

func (g *Game) IsPotentialGameForSet(search *Set) bool {
	possibleGame := true
	for _, set := range g.Sets {
		if set.Red > search.Red || set.Green > search.Green || set.Blue > search.Blue {
			possibleGame = false
			break
		}
	}
	return possibleGame
}

func (g *Game) LowestPossibleDice() Set {
	minimumSet := Set{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
	for _, set := range g.Sets {
		if set.Red > minimumSet.Red {
			minimumSet.Red = set.Red
		}
		if set.Blue > minimumSet.Blue {
			minimumSet.Blue = set.Blue
		}
		if set.Green > minimumSet.Green {
			minimumSet.Green = set.Green
		}
	}
	return minimumSet
}

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

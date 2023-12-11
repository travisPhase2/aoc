package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const ( // types of hands
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var part1CardValues = map[byte]int {
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var part2CardValues = map[byte]int {
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type Game struct { 
	Cards string
	Bid int
	Hand int
}

func NewGame(cards string, bid int) *Game {
	return &Game{
		Cards: cards,
		Bid: bid,
		Hand: calculateHand(cards),
	}
}

func calculateHand(cards string) int {
	countMap := map[rune]int{}
	for _, card := range cards {
		countMap[card]++
	}

	// do stuff with Jokers here
	applyJokers(countMap)

	var counts []int
	for _, count := range countMap { 
		counts = append(counts, count)
	}

	slices.Sort(counts)
	slices.Reverse(counts)

	result := 0
	switch counts[0] {
	case 5:
		result = FiveOfAKind
	case 4:
		result = FourOfAKind
	case 3:
		if counts[1] == 2 {
			result = FullHouse
		} else { 
			result = ThreeOfAKind
		}
	case 2:
		if counts[1] == 2 {
			result = TwoPair
		} else { 
			result = OnePair
		}
	case 1:
		result = HighCard
	}
	return result
}

func applyJokers(countMap map[rune]int) {
	for countMap['J'] > 0 {
		maxCardCount := -1
		var cardToIncrement rune
		for key, count := range countMap {
			if key == 'J' {
				continue
			}
			if count > maxCardCount {
				maxCardCount = count
				cardToIncrement = key
			}
		}
		countMap['J']--
		countMap[cardToIncrement]++
	}
}

func parseInput() []*Game {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	var games []*Game
	for scanner.Scan() { 
		line := strings.Fields(scanner.Text())
		cards := line[0]
		bid, _ := strconv.Atoi(line[1])
		games = append(games, NewGame(cards, bid))
	}

	return games
}

func main() {
	games := parseInput()

	// part 1 sort
	// sort.Slice(games, func(i, j int) bool {
	// 	// first sort by hand value
	// 	firstGame, secondGame := *games[i], *games[j]
	// 	if (firstGame.Hand < secondGame.Hand) {
	// 		return true
	// 	}

	// 	if (games[i].Hand > games[j].Hand) {
	// 		return false
	// 	}

	// 	// sort by card values if the hand values are the same
	// 	for c := 0; c < 5; c++ {
	// 		c1, c2 := part1CardValues[firstGame.Cards[c]], part1CardValues[secondGame.Cards[c]]

	// 		if c1 < c2 { 
	// 			return true
	// 		}

	// 		if c1 > c2 { 
	// 			return false
	// 		}
	// 	}

	// 	return false
	// })

	sort.Slice(games, func(i, j int) bool {
		// first sort by hand value
		firstGame, secondGame := *games[i], *games[j]
		if (firstGame.Hand < secondGame.Hand) {
			return true
		}

		if (games[i].Hand > games[j].Hand) {
			return false
		}

		// sort by card values if the hand values are the same
		for c := 0; c < 5; c++ {
			c1, c2 := part2CardValues[firstGame.Cards[c]], part2CardValues[secondGame.Cards[c]]

			if c1 < c2 { 
				return true
			}

			if c1 > c2 { 
				return false
			}
		}

		return false
	})

	totalWinnings := 0
	for idx, g := range games {
		totalWinnings += g.Bid * (idx + 1)
	}
	fmt.Println(totalWinnings)
}
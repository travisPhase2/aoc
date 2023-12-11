package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	Left, Right string
}

func parseInput() ([]string, map[string]Pair) { 
	f, _ := os.Open("input")
	s := bufio.NewScanner(f)

	s.Scan()
	instructions := strings.Split(s.Text(), "")
	network := map[string]Pair{}
	for s.Scan() { 
		if strings.Trim(s.Text(), " ") == "" {
			continue
		}
		parts := strings.Fields(s.Text())
		key := parts[0]
		left := strings.TrimRight(strings.TrimLeft(parts[2], "("), ",")
		right := strings.TrimRight(parts[3], ")")
		network[key] = Pair{Left: left, Right: right}
	}

	return instructions, network
}

func solvePart1(instructions []string, network map[string]Pair, start string) int { 
	continueSearch := true
	currentInstruction := 0
	steps := 0
	for continueSearch { 
		inst := instructions[currentInstruction]
		pair := network[start]
		switch inst {
		case "R":
			if pair.Right == "ZZZ" {
				steps += 1
				continueSearch = false
			} else { 
				steps += 1
				if currentInstruction == len(instructions) - 1 {
					currentInstruction = 0
				} else {
					currentInstruction += 1
				}
				start = pair.Right
			}
		case "L":
			if pair.Left == "ZZZ" {
				steps += 1
				continueSearch = false
			} else { 
				steps += 1
				if currentInstruction == len(instructions) - 1 {
					currentInstruction = 0
				} else {
					currentInstruction += 1
				}
				start = pair.Left
			}
		}
	}

	return steps
}

func solvePart2(instructions []string, network map[string]Pair, start string) int { 
	continueSearch := true
	currentInstruction := 0
	steps := 0
	for continueSearch { 
		inst := instructions[currentInstruction]
		pair := network[start]
		switch inst {
		case "R":
			if strings.HasSuffix(pair.Right, "Z") {
				steps += 1
				continueSearch = false
			} else { 
				steps += 1
				if currentInstruction == len(instructions) - 1 {
					currentInstruction = 0
				} else {
					currentInstruction += 1
				}
				start = pair.Right
			}
		case "L":
			if strings.HasSuffix(pair.Left, "Z") {
				steps += 1
				continueSearch = false
			} else { 
				steps += 1
				if currentInstruction == len(instructions) - 1 {
					currentInstruction = 0
				} else {
					currentInstruction += 1
				}
				start = pair.Left
			}
		}
	}

	return steps
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	instructions, network := parseInput()

	var startingPositions []string
	for key := range network { 
		if strings.HasSuffix(key, "A") {
			startingPositions = append(startingPositions, key)
		}
	}

	var lengths []int
	for _, pos := range startingPositions {
		lengths = append(lengths, solvePart2(instructions, network, pos))
	}

	fmt.Println(lengths)

	fmt.Println(LCM(lengths[0], lengths[1], lengths...))
}

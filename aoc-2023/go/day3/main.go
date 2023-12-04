package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, _ := os.ReadFile("input-sample.txt")

	grid := map[image.Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r != '.' && !unicode.IsDigit(r) {
				grid[image.Point{x, y}] = r
			}
		}
	}

	directions := []image.Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	part1, part2 := 0, 0
	// [(x, y): [1, 2, 3...], ...]
	points := map[image.Point][]int{}
	for y, row := range strings.Fields(string(input)) {
		for _, found := range regexp.MustCompile(`\d+`).FindAllStringIndex(row, -1) {
			bounds := map[image.Point]struct{}{}
			start, end := found[0], found[1]

			for x := start; x < end; x++ {
				for _, point := range directions {
					bounds[image.Point{x, y}.Add(point)] = struct{}{}
				}
			}

			number, _ := strconv.Atoi(row[start:end])
			for b := range bounds {
				if _, ok := grid[b]; ok {
					points[b] = append(points[b], number)
					part1 += number
				}
			}
		}
	}

	for p, hits := range points {
		if grid[p] == '*' && len(hits) == 2 {
			part2 += hits[0] * hits[1]
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

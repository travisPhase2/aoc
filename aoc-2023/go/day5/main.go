package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type GardenMap struct {
	Name   string
	Ranges []Range
}

type Range struct {
	DestStart   int
	SourceStart int
	RangeLen    int
}

func (m *GardenMap) GetLocation(search int) int {
	for _, r := range m.Ranges {
		if search >= r.SourceStart && search <= r.SourceStart+r.RangeLen {
			return r.DestStart + (search - r.SourceStart)
		}
	}
	return search
}

func (m *GardenMap) GetLocationInverse(search int) int {
	for _, r := range m.Ranges {
		if search >= r.DestStart && search < r.DestStart+r.RangeLen {
			return r.SourceStart + (search - r.DestStart)
		}
	}
	return search
}

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
		log.Fatal("could not read file")
	}

	lines := strings.Split(string(data), "\n")

	var trimmedLines []string
	for _, line := range lines {
		trimmedLines = append(trimmedLines, strings.Trim(line, " "))
	}

	return trimmedLines
}

func parseMaps(lines []string) []GardenMap {
	var maps []GardenMap
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		input := strings.Split(line, " ")
		var group []string
		if len(input) == 2 {
			linesToGrab := 0
			for j := i + 1; j < len(lines); j++ {
				lineToInspect := strings.Split(lines[j], " ")
				if len(lineToInspect) == 3 {
					linesToGrab += 1
				} else if len(lineToInspect) == 2 {
					break
				}
			}
			group = append(group, input[0])
			for j := 1; j <= linesToGrab; j++ {
				group = append(group, lines[i+j])
			}
		}
		if len(group) != 0 {
			groupName := group[0]
			gardenMap := GardenMap{
				Name: groupName,
			}
			for _, r := range group[1:] {
				split := strings.Split(r, " ")
				destStart, _ := strconv.Atoi(split[0])
				sourceStart, _ := strconv.Atoi(split[1])
				rangeLen, _ := strconv.Atoi(split[2])

				gardenMap.Ranges = append(gardenMap.Ranges, Range{
					DestStart:   destStart,
					SourceStart: sourceStart,
					RangeLen:    rangeLen,
				})
			}
			maps = append(maps, gardenMap)
		}
	}
	return maps
}

func parseSeeds(line string) []int {
	var seeds []int
	seedNumsStr := strings.Split(line, ":")[1]
	seedNums := filter(strings.Split(seedNumsStr, " "), func(s string) bool {
		return strings.Trim(s, " ") != ""
	})
	for _, num := range seedNums {
		n, _ := strconv.Atoi(num)
		seeds = append(seeds, n)
	}
	return seeds
}

func main() {
	lines := readInputToLines("input")
	seeds := parseSeeds(lines[0])
	maps := parseMaps(lines[1:])

	var locations []int
	for _, s := range seeds {
		search := s
		for _, m := range maps {
			search = m.GetLocation(search)
		}
		locations = append(locations, search)
	}

	// min := slices.Min(locations)
	// fmt.Println("\n part 1: ", min)

	min := 0
	for loc := 0; ; loc++ {
		x := loc
		for m := len(maps) - 1; m >= 0; m-- {
			x = maps[m].GetLocationInverse(x)
		}

		for s := 0; s < len(seeds); s += 2 {
			if x >= seeds[s] && x <= seeds[s]+seeds[s+1] {
				min = loc
				break
			}
		}
		break
	}

	fmt.Println(min)
}

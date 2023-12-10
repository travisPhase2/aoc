package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() { 
	// times, distances := parseInputPart1()

	// possibilities := 1
	// for idx, raceTime := range times {
	// 	count := 0 
	// 	for t := 0; t < raceTime; t++ {
	// 		if (raceTime-t)*t > distances[idx] {
	// 			count++
	// 		}
	// 	}

	// 	possibilities *= count
	// }

	// fmt.Println(possibilities)

	time, distance := parseInputPart2()

	count := 0 
	for t := 0; t < time; t++ {
		if (time-t)*t > distance {
			count++
		}
	}

	fmt.Println(count)
}

func parseInputPart1() ([]int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var times []int
	for _, t := range strings.Fields(scanner.Text())[1:] {
		n, _ := strconv.Atoi(t)
		times = append(times, n)
	}

	scanner.Scan()
	var distances []int
	for _, d := range strings.Fields(scanner.Text())[1:] {
		n, _ := strconv.Atoi(d)
		distances = append(distances, n)
	}
	return times, distances
}

func parseInputPart2() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	time, _ := strconv.Atoi(strings.Join(strings.Fields(scanner.Text())[1:], ""))
	scanner.Scan()
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(scanner.Text())[1:], ""))

	return time, distance
}
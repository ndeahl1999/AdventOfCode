package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]

	LoadData(fileName)
	fmt.Printf("Part 1 answer is: %d\n", solvePart1())
	fmt.Printf("Part 2 answer is %d\n", solvePart2())

}

var equations map[int][]int

func LoadData(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	equations = make(map[int][]int, 0)
	fs := bufio.NewScanner(file)

	for fs.Scan() {
		line := strings.Split(fs.Text(), ":")
		result, _ := strconv.Atoi(line[0])
		nums := strings.Split(line[1], " ")[1:]
		numList := make([]int, 0)
		for _, numStr := range nums {
			numVal, _ := strconv.Atoi(numStr)
			numList = append(numList, numVal)
		}

		equations[result] = numList

	}

}

func solvePart1() int {
	total := 0

	for result, vals := range equations {
		if canBeTrue(result, vals, false) {
			total += result
		}
	}

	return total
}

func canBeTrue(target int, vals []int, isPart2 bool) bool {
	if len(vals) == 1 {
		return target == vals[0]
	}

	allTotals := make(map[int][]int, 0)
	allTotals[0] = []int{vals[0]}

	for index, val := range vals {
		if index == 0 {
			continue
		}

		localTotals := make([]int, 0)
		for _, prevTotal := range allTotals[index-1] {
			if prevTotal+val <= target {
				localTotals = append(localTotals, prevTotal+val)
			}
			if prevTotal*val <= target {
				localTotals = append(localTotals, prevTotal*val)
			}

			if isPart2 {
				concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", prevTotal, val))
				if concat <= target {
					localTotals = append(localTotals, concat)
				}
			}
		}

		if len(localTotals) == 0 {
			break
		}
		allTotals[index] = localTotals
	}
	return slices.Contains(allTotals[len(vals)-1], target)
}

func solvePart2() int {
	total := 0

	for result, vals := range equations {
		if canBeTrue(result, vals, true) {
			total += result
		}
	}

	return total
}

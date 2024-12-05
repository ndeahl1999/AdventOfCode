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

	answer1, answer2 := solvePart1and2(fileName)
	fmt.Printf("Part 1 answer is: %d\nPart 2 answer is: %d\n", answer1, answer2)

}

func solvePart1and2(fileName string) (int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)

	rules := make(map[int][]int)
	updateList := make([][]int, 0)

	readUpdates := false
	for fs.Scan() {
		if fs.Text() == "" {
			readUpdates = true
			continue
		}

		if !readUpdates {
			line := strings.Split(fs.Text(), "|")
			l1, _ := strconv.Atoi(line[0])
			l2, _ := strconv.Atoi(line[1])
			rules[l1] = append(rules[l1], l2)
		} else {
			line := strings.Split(fs.Text(), ",")
			updateLine := make([]int, 0)
			for _, update := range line {
				intVal, _ := strconv.Atoi(update)
				updateLine = append(updateLine, intVal)
			}
			updateList = append(updateList, updateLine)
		}

	}

	total1, total2 := 0, 0

	for _, updates := range updateList {
		isValid := true
		for pageIndex, pageNum := range updates {
			if pageIndex == 0 {
				continue
			}
			for i := 0; i < pageIndex; i++ {
				if slices.Contains(rules[pageNum], updates[i]) {
					isValid = false
					break
				}
			}
			if isValid == false {
				break
			}
		}

		if isValid {
			total1 += updates[int(len(updates)/2)]
		} else {
			sorted := sortPages(updates, rules)
			total2 += sorted[int(len(sorted)/2)]
		}

	}

	return total1, total2
}

func sortPages(pages []int, rules map[int][]int) []int {
	slices.SortFunc(pages, func(a, b int) int {
		if slices.Contains(rules[a], b) {
			return -1
		}
		return 1
	})

	return pages
}

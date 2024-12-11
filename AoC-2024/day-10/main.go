package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	fileName := os.Args[1]

	LoadData(fileName)
	fmt.Printf("Part 1 answer is: %d\n", solvePart1())
	fmt.Printf("Part 2 answer is %d\n", solvePart2())

}

type Pos struct{ i, j int }

var grid [][]int

func LoadData(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)

	grid = make([][]int, 0)
	for fs.Scan() {
		row := make([]int, 0)
		for _, char := range fs.Text() {
			val, _ := strconv.Atoi(string(char))
			row = append(row, val)
		}
		grid = append(grid, row)
	}

}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%d", val)
		}
		fmt.Print("\n")
	}
}

var trailHeads map[Pos][]Pos

func computeScores() int {
	total := 0
	for _, headTotal := range trailHeads {
		total += len(headTotal)
	}

	return total
}

func solvePart1() int {
	trailHeads = make(map[Pos][]Pos)
	for i, row := range grid {
		for j, val := range row {
			if val == 0 {
				searchGrid(Pos{i, j}, Pos{i, j}, 0, true)
			}
		}
	}

	return computeScores()
}

func searchGrid(startPos, currPos Pos, desiredNum int, part1 bool) {
	if currPos.i < 0 || currPos.i >= len(grid) || currPos.j < 0 || currPos.j >= len(grid[0]) {
		return
	}

	if grid[currPos.i][currPos.j] != desiredNum {
		return
	}

	if grid[currPos.i][currPos.j] == 9 && desiredNum == 9 {
		if part1 {
			if !slices.Contains(trailHeads[startPos], currPos) {
				trailHeads[startPos] = append(trailHeads[startPos], currPos)
			}
		} else {
			trailHeads[startPos] = append(trailHeads[startPos], currPos)
		}
		return
	}
	newNum := desiredNum + 1
	searchGrid(startPos, Pos{currPos.i - 1, currPos.j}, newNum, part1)
	searchGrid(startPos, Pos{currPos.i, currPos.j + 1}, newNum, part1)
	searchGrid(startPos, Pos{currPos.i + 1, currPos.j}, newNum, part1)
	searchGrid(startPos, Pos{currPos.i, currPos.j - 1}, newNum, part1)

}

func solvePart2() int {
	trailHeads = make(map[Pos][]Pos)
	for i, row := range grid {
		for j, val := range row {
			if val == 0 {
				searchGrid(Pos{i, j}, Pos{i, j}, 0, false)
			}
		}
	}

	return computeScores()
}

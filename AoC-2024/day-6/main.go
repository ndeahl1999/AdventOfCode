package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"
)

func main() {
	fileName := os.Args[1]

	fmt.Printf("Part 1 answer is: %d\n", solvePart1(fileName))
	fmt.Printf("Part 2 answer is %d\n", solvePart2())

}

type Coords struct{ x, y int }

type Position struct {
	isObstacle bool
	visited    bool
}

var grid [][]Position
var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var startI, startJ int

func solvePart1(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)

	i := 0
	for fs.Scan() {
		var row []Position
		for j, char := range fs.Text() {

			row = append(row, Position{char == '#', false})

			if char == '^' {
				startI = i
				startJ = j
			}
		}
		grid = append(grid, row)
		i += 1
	}

	i, j := startI, startJ

	currDir := 0

	for {
		grid[i][j].visited = true
		newI, newJ := i+directions[currDir][0], j+directions[currDir][1]
		if newI < 0 || newJ < 0 || newI >= len(grid) || newJ >= len(grid[0]) {
			break
		}
		if grid[newI][newJ].isObstacle {
			currDir = (currDir + 1) % 4

		} else {
			i, j = newI, newJ
		}

	}

	total := 0
	for _, row := range grid {
		for _, pos := range row {
			if pos.visited {
				total += 1
			}
		}
	}

	return total
}

func solvePart2() int {
	totalLoops := 0
	grid[startI][startJ].visited = false
	wg := sync.WaitGroup{}
	for i, row := range grid {
		for j, pos := range row {
			if pos.visited {
				wg.Add(1)
				go func() {
					defer wg.Done()
					if checkIsLoop(i, j) {
						totalLoops += 1
					}
				}()
				wg.Wait()
			}
		}
	}

	return totalLoops
}

func checkIsLoop(obstacleI, obstacleJ int) bool {
	visit := make(map[Coords][]int)

	i, j := startI, startJ
	currDir := 0

	for {
		newI, newJ := i+directions[currDir][0], j+directions[currDir][1]
		if newI < 0 || newJ < 0 || newI >= len(grid) || newJ >= len(grid[0]) {
			break
		}
		if (newI == obstacleI && newJ == obstacleJ) || grid[newI][newJ].isObstacle {
			currDir = (currDir + 1) % 4
		} else {
			if slices.Contains(visit[Coords{newI, newJ}], currDir) {
				return true
			}
			visit[Coords{newI, newJ}] = append(visit[Coords{newI, newJ}], currDir)

			i, j = newI, newJ

		}

	}

	return false
}

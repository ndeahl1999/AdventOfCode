package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]

	LoadData(fileName)
	fmt.Printf("Part 1 answer is: %d\n", solvePart1())
	fmt.Printf("Part 2 answer is %d\n", solvePart2())

}

type Coords struct{ x, y int }

var grid [][]rune
var antinodes map[rune][]Coords

func LoadData(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid = make([][]rune, 0)
	antinodes = make(map[rune][]Coords)
	fs := bufio.NewScanner(file)

	i := 0
	for fs.Scan() {
		grid = append(grid, []rune(fs.Text()))

		for j, char := range fs.Text() {
			if char != '.' {
				antinodes[char] = append(antinodes[char], Coords{i, j})
			}
		}

		i++
	}

}

func solvePart1() int {

	uniqueAntinodes := make(map[Coords]struct{})
	for _, antinodeCoords := range antinodes {
		for i := range antinodeCoords {
			for j := range antinodeCoords {
				if i == j {
					continue
				}

				dx := antinodeCoords[j].x - antinodeCoords[i].x
				dy := antinodeCoords[j].y - antinodeCoords[i].y

				antX := antinodeCoords[i].x + (2 * dx)
				antY := antinodeCoords[i].y + (2 * dy)

				if antX >= 0 && antX < len(grid) && antY >= 0 && antY < len(grid[0]) {
					uniqueAntinodes[Coords{antX, antY}] = struct{}{}
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

func solvePart2() int {
	diagLength := len(grid) + 1/2
	uniqueAntinodes := make(map[Coords]struct{})
	for _, antinodeCoords := range antinodes {
		for i := range antinodeCoords {
			for j := range antinodeCoords {
				if i == j {
					continue
				}

				dx := antinodeCoords[j].x - antinodeCoords[i].x
				dy := antinodeCoords[j].y - antinodeCoords[i].y

				for dist := -diagLength; dist <= diagLength; dist++ {
					antX := antinodeCoords[i].x + (dist * dx)
					antY := antinodeCoords[i].y + (dist * dy)

					if antX >= 0 && antX < len(grid) && antY >= 0 && antY < len(grid[0]) {
						uniqueAntinodes[Coords{antX, antY}] = struct{}{}
					}
				}

			}
		}
	}

	return len(uniqueAntinodes)
}

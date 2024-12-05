package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]

	fmt.Printf("Part 1 answer is: %d\n", solvePart1(fileName))
	fmt.Printf("Part 2 answer is: %d\n", solvePart2(fileName))
}

func solvePart1(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [][]rune

	fs := bufio.NewScanner(file)
	total := 0
	for fs.Scan() {
		matrix = append(matrix, []rune(fs.Text()))
	}

	for i, row := range matrix {
		for j, cell := range row {
			if cell == 'X' {
				total += checkDiags(&matrix, i, j)
			}
		}
	}

	return total
}

func checkDiags(matrix *[][]rune, startI, startJ int) int {
	total := 0

	directions := [][]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	nextChars := []rune{'M', 'A', 'S'}

	for _, direction := range directions {
		i, j := startI, startJ
		isValid := true
		for _, char := range nextChars {
			i = i + direction[0]
			j = j + direction[1]

			if i < 0 || i >= len(*matrix) || j < 0 || j >= len((*matrix)[0]) {
				isValid = false
				break
			}

			if (*matrix)[i][j] != char {
				isValid = false
				break
			}

		}

		if isValid {
			total += 1
		}
	}

	return total
}

func solvePart2(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [][]rune

	fs := bufio.NewScanner(file)
	total := 0
	for fs.Scan() {
		matrix = append(matrix, []rune(fs.Text()))
	}

	for i, row := range matrix {
		for j, cell := range row {
			if cell == 'A' {
				if checkMASDiags(&matrix, i, j) {
					total += 1
				}
			}
		}
	}

	return total
}

func checkMASDiags(matrix *[][]rune, i, j int) bool {
	diags := [][][]int{
		{{-1, -1}, {1, 1}},
		{{-1, 1}, {1, -1}},
	}

	if i-1 < 0 || i+1 >= len(*matrix) || j-1 < 0 || j+1 >= len((*matrix)[0]) {
		return false
	}

	for _, diag := range diags {
		charTop := (*matrix)[i+diag[0][0]][j+diag[0][1]]
		charBot := (*matrix)[i+diag[1][0]][j+diag[1][1]]

		if (charTop == 'M' && charBot == 'S') || (charTop == 'S' && charBot == 'M') {
			continue
		}
		return false
	}

	return true

}

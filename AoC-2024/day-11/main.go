package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]

	LoadData(fileName)
	fmt.Printf("Part 1 answer is: %d\n", solvePart1())
	fmt.Printf("Part 2 answer is %d\n", solvePart2())

}

var startingStones map[int]int

func LoadData(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	startingStones = make(map[int]int)
	fs := bufio.NewScanner(file)
	for fs.Scan() {
		for _, char := range strings.Split(fs.Text(), " ") {
			val, _ := strconv.Atoi(string(char))
			startingStones[val]++
		}
	}

	fmt.Println(startingStones)

}

func countStones(stones map[int]int) int {
	total := 0
	for _, count := range stones {
		total += count
	}

	return total
}

func solvePart1() int {
	stones := startingStones
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	return countStones(stones)
}

func blink(stones map[int]int) map[int]int {
	updatedStones := make(map[int]int)
	for num, count := range stones {
		if num == 0 {
			updatedStones[1] += count
		} else if countDigits(num)%2 == 0 {
			l, r := splitNum(num)
			updatedStones[l] += count
			updatedStones[r] += count
		} else {
			updatedStones[num*2024] += count
		}

	}

	return updatedStones
}

func countDigits(num int) int {
	strNum := strconv.Itoa(num)
	return len(strNum)
}

func splitNum(num int) (int, int) {
	strNum := strconv.Itoa(num)
	len := len(strNum) / 2
	num1, _ := strconv.Atoi(strNum[:len])
	num2, _ := strconv.Atoi(strNum[len:])
	return num1, num2

}

func solvePart2() int {
	stones := startingStones
	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}

	return countStones(stones)
}

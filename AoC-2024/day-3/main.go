package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fileName := os.Args[1]

	fmt.Printf("Total of multiplications is %d\nTotal with instructions is %d\n", solvePart1(fileName), solvePart2(fileName))
}

func solvePart1(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	total := 0
	for fs.Scan() {

		re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)

		matches := re.FindAllString(fs.Text(), -1)
		for _, match := range matches {
			total += multiplyVals(match)
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

	fs := bufio.NewScanner(file)
	total := 0
	multEnabled := true
	for fs.Scan() {
		re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don\'t\(\)`)

		ops := re.FindAllString(fs.Text(), -1)
		for _, op := range ops {
			if op == "do()" {
				multEnabled = true
				continue
			} else if op == "don't()" {
				multEnabled = false
				continue
			}

			if multEnabled {
				total += multiplyVals(op)
			}

		}
	}
	return total
}

func multiplyVals(match string) int {
	number_re := regexp.MustCompile(`[0-9]{1,3}`)
	nums := number_re.FindAllString(match, 2)
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	return num1 * num2
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalSafe := 0
	totalSafeDampened := 0
	fs := bufio.NewScanner(file)
	for fs.Scan() {
		line := strings.Split(fs.Text(), " ")

		level := make([]int, len(line))
		for i := range line {
			val, _ := strconv.Atoi(line[i])
			level[i] = val
		}

		if isReportSafe(level) {
			totalSafe++
			totalSafeDampened++
		} else if isReportSafeWithDampener(level) {
			totalSafeDampened++
		}
	}

	fmt.Printf("Total safe is %d\n", totalSafe)
	fmt.Printf("Total safe with dampener is %d\n", totalSafeDampened)
}

func isReportSafe(level []int) bool {
	isIncreasing := level[0] < level[len(level)-1]
	isSafe := true
	for i := 1; i < len(level); i++ {

		if (isIncreasing && level[i] < level[i-1]) || (!isIncreasing && level[i] > level[i-1]) {
			isSafe = false
			continue
		}

		dist := int(math.Abs(float64(level[i] - level[i-1])))
		if dist < 1 || dist > 3 {
			isSafe = false
		}
	}
	return isSafe
}

func isReportSafeWithDampener(level []int) bool {
	for i := 0; i < len(level); i++ {
		copyLevel := make([]int, len(level))
		copy(copyLevel, level)
		copyLevel = append(copyLevel[:i], copyLevel[i+1:]...)
		if isReportSafe(copyLevel) {
			return true
		}
	}
	return false
}

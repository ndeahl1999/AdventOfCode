package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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

	fs := bufio.NewScanner(file)

	var lArr, rArr []int
	var totalDist float64
	for fs.Scan() {
		if err := fs.Err(); err != nil {
			log.Fatal(err)
		}

		split := strings.Split(fs.Text(), "   ")
		lInt, _ := strconv.Atoi(split[0])
		rInt, _ := strconv.Atoi(split[1])

		lArr = append(lArr, lInt)
		rArr = append(rArr, rInt)

	}

	slices.Sort(lArr)
	slices.Sort(rArr)

	for i := range lArr {
		totalDist += math.Abs(float64(lArr[i] - rArr[i]))
	}

	fmt.Printf("Total Dist is: %d\n", int(totalDist))

	totalSim := 0
	freqMap := make(map[int]int, 0)
	for i := range rArr {
		freqMap[rArr[i]]++
	}

	for i := range lArr {
		totalSim += lArr[i] * freqMap[lArr[i]]
	}

	fmt.Printf("Similarity score is: %d\n", totalSim)
}

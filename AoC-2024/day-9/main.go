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

type MemBlock struct{ ID, len int }

var disk []MemBlock

func LoadData(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)

	id := 0
	for fs.Scan() {
		for i, char := range fs.Text() {
			len, _ := strconv.Atoi(string(char))
			if i%2 == 0 {
				disk = append(disk, MemBlock{id, len})
				id++
			} else {
				disk = append(disk, MemBlock{-1, len})
			}

		}
	}

}

func printMem(printDisk []MemBlock) {
	for _, mem := range printDisk {
		for i := 0; i < mem.len; i++ {
			if mem.ID == -1 {
				fmt.Print(".")
			} else {

				fmt.Print(mem.ID)
			}
		}

	}

	fmt.Print("\n")
}

func calculateChecksum(blocks []MemBlock) int {
	checkSum := 0

	pos := 0
	for _, block := range blocks {
		for i := 0; i < block.len; i++ {
			if block.ID != -1 {
				checkSum += pos * block.ID

			}
			pos++
		}
	}

	return checkSum
}

func solvePart1() int {
	diskCopy := make([]MemBlock, len(disk))
	copy(diskCopy, disk)

	l, r := 1, len(diskCopy)-1
	for l <= r {

		blockL := diskCopy[l]
		if blockL.len == diskCopy[r].len {
			diskCopy[l] = diskCopy[r]
			diskCopy[r] = blockL
			l += 2
			r -= 2
		} else if blockL.len > diskCopy[r].len {
			diskCopy = slices.Insert(diskCopy, l, diskCopy[r])
			diskCopy[l+1].len -= diskCopy[r+1].len
			diskCopy[r+1].ID = -1
			l += 1
			r -= 1
		} else {
			diskCopy[r].len -= blockL.len
			diskCopy[l].ID = diskCopy[r].ID
			l += 2
		}

	}

	return calculateChecksum(diskCopy)
}

func solvePart2() int {
	diskCopy := make([]MemBlock, len(disk))
	copy(diskCopy, disk)

	r := len(diskCopy) - 1
	for r > 0 {
		for r > 0 && diskCopy[r].ID == -1 {
			r--
		}
		if r == 0 {
			break
		}

		l := 1
		for l < r && !(diskCopy[l].ID == -1 && diskCopy[l].len >= diskCopy[r].len) {
			l++
		}

		if diskCopy[l].len == diskCopy[r].len {
			blockL := diskCopy[l]
			diskCopy[l] = diskCopy[r]
			diskCopy[r] = blockL
		} else {
			diskCopy = slices.Insert(diskCopy, l, diskCopy[r])
			diskCopy[l+1].len -= diskCopy[r+1].len
			diskCopy[r+1].ID = -1
		}

		r--
	}

	return calculateChecksum(diskCopy)
}

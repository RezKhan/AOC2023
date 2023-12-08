package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getArrayOfMap(searchstr string, lines []string) [][]int {
	start := slices.Index(lines, searchstr) + 1
	var tempArray [][]int
	for _, line := range lines[start:] {
		var nums []int
		snums := strings.Split(line, " ")
		for _, num := range snums {
			t, err := strconv.Atoi(num)
			if err == nil {
				nums = append(nums, t)
			}
		}
		tempArray = append(tempArray, nums)
		// log.Println(len(line), line)
		if strings.TrimSpace(line) == "" {
			break
		}
	}
	return tempArray
}

func partOne(lines []string) {
	// MAP ARRAY
	// ARR[0] to ARR[0+2] = DESTINATION RANGE
	// ARR [1] to ARR[1+2] = SOURCE RANGE
	// IF SEED OR START POS < ARR[0] THEN DESTINATION = SEED OR START POS
	// START WITH SEED VALUE, CHECK FOR SOURCE RANGE THEN DEPOSIT AT
	// SAME RELATIVE POSITION FROM THE DESTINATION RANGE

	seeds := strings.Split(lines[0][7:], " ")

	seedsToSoilMapArray := getArrayOfMap("seed-to-soil map:", lines)
	log.Println(seeds)
	log.Println(seedsToSoilMapArray)
}

func ReadMaps(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	var lines []string
	for fScanner.Scan() {
		line := fScanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func main() {
	filePath := "./d05test.txt"
	lines := ReadMaps(filePath)
	partOne(lines)
}

package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkWinners(winningNums []int, cardNums []int) (int, int) {
	counter := 0
	for i := 0; i < len(cardNums); i++ {
		if slices.Contains(winningNums, cardNums[i]) {
			counter++
		}
	}
	if counter > 0 {
		return int(math.Pow(2, float64(counter-1))), counter
	} else {
		return counter, counter
	}
}

func parseLine(line string) ([]int, []int) {
	strToNum := func(str []string) []int {
		var intArray []int
		for _, nums := range str {
			if n, err := strconv.Atoi(nums); err == nil {
				intArray = append(intArray, n)
			}
		}
		return intArray
	}
	colonPos := strings.Index(line, ":")
	pipePos := strings.Index(line, "|")
	winningNumsList := strings.Split(line[colonPos+1:pipePos], " ")
	cardNumsList := strings.Split(line[pipePos+1:], " ")
	winningNums := strToNum(winningNumsList)
	cardNums := strToNum(cardNumsList)

	return winningNums, cardNums
}

func ReadCards(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)

	var lines []string
	var cardwins []int
	for fScanner.Scan() {
		line := fScanner.Text()
		lines = append(lines, line)
		cardwins = append(cardwins, 0) // populates an empty array of wins
	}

	winnings := 0
	// winnings2 := 0
	// Part 1:
	for i := 0; i < len(lines); i++ {
		winningNums, cardNums := parseLine(lines[i])
		win, _ := checkWinners(winningNums, cardNums)
		winnings += win
	}
	log.Println("Part 1, winnings: ", winnings)
	winnings = 0

	// Part 2:
	for i := 0; i < len(lines); i++ {
		counter := 0
		for j := 0; j <= cardwins[i]; j++ {
			winningNums, cardNums := parseLine(lines[i])
			win, count := checkWinners(winningNums, cardNums)
			winnings += win
			counter = count
			n := 0
			for n = 1; n <= counter; n++ {
				cardwins[i+n]++
			}
		}
	}
	sumrange := 0
	for i := 0; i < len(cardwins); i++ {
		sumrange += cardwins[i]
	}
	log.Println("Part 2, count of cards: ", sumrange+len(lines))
}

func main() {
	// filePath := "./d04test.txt"
	filePath := "./d04input.txt"
	ReadCards(filePath)
}

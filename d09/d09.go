package main

import (
	"log"
	rf "readfile"
	"slices"
	"strconv"
	"strings"
)

func subHistory(history []int) []int {
	var subHist []int
	lh := len(history)
	for i := 1; i < lh; i++ {
		n := history[i] - history[i-1]
		subHist = append(subHist, n)
	}
	allZeros := true
	for _, n := range subHist {
		if n == 0 {
			continue
		} else {
			allZeros = false
		}
	}
	if !allZeros {
		subHist = subHistory(subHist)
	}
	history = append(history, history[lh-1]+subHist[len(subHist)-1])
	return history
}

func PartTwo(lines []string) {
	var histories [][]int
	for _, line := range lines {
		ns := strings.Split(line, " ")
		var history []int
		for _, n := range ns {
			h, err := strconv.Atoi(n)
			if err == nil {
				history = append(history, h)
			}
		}
		slices.Reverse(history)
		histories = append(histories, history)
	}
	sumh := 0
	for _, history := range histories {
		history := subHistory(history)
		sumh += history[len(history)-1]
	}
	log.Println("Sum of last numbers in history: ", sumh)
}


func PartOne(lines []string) {
	var histories [][]int
	for _, line := range lines {
		ns := strings.Split(line, " ")
		var history []int
		for _, n := range ns {
			h, err := strconv.Atoi(n)
			if err == nil {
				history = append(history, h)
			}
		}
		histories = append(histories, history)
	}
	sumh := 0
	for _, history := range histories {
		history := subHistory(history)
		sumh += history[len(history)-1]
	}
	log.Println("Sum of last numbers in history: ", sumh)
}

func main() {
	// filePath := "d09test.txt"
	filePath := "d09input.txt"

	lines := rf.ReadFile(filePath)
	// PartOne(lines)
	PartTwo(lines)
}

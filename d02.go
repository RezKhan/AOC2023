package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func diceCountSet(set string, targetDice map[string]int) bool {
	gameDice := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	i := 0
	for i < len(set) {
		var chunk string
		next := 0

		comma := strings.Index(set[i:], ",")
		if comma < 0 {
			comma = i
		}
		next = i + comma
		if next == -1 || next > len(set) {
			next = len(set)
		}

		chunk = set[i:next]
		c := strings.Split(chunk, " ")
		n, _ := strconv.Atoi(c[0])
		if len(c) > 1 {
			gameDice[c[1]] += n
		}
		i = next + 2
	}
	// log.Println("target: ", targetDice, "game: ", gameDice)
	if targetDice["red"] >= gameDice["red"] && targetDice["green"] >= gameDice["green"] && targetDice["blue"] >= gameDice["blue"] {
		return true
	} else {
		return false
	}
}

func sumSlice(n []int) int {
	result := 0
	for i := 0; i < len(n); i++ {
		result += n[i]
	}
	return result
}

func ReadFile(filePath string) {
	targetDice := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	var ids []int
	n := 1
	for fScanner.Scan() {
		game := (fScanner.Text()[strings.Index(fScanner.Text(), ":")+2:])
		sets := strings.Split(game, "; ")
		setsTested := true
		for i := 0; i < len(sets); i++ {
			if !diceCountSet(sets[i], targetDice) {
				setsTested = false
			}
		}
		if setsTested {
			ids = append(ids, n)
		}
		n++
	}
	log.Println(sumSlice(ids))
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Enter filename of test file")
	}

	filePath := os.Args[1]
	ReadFile(filePath)
}

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func findChunk(str string, i int) (string, int) {
	next := 0

	comma := strings.Index(str[i:], ",")
	if comma < 0 {
		comma = len(str)
	}

	next = i + comma
	if next == -1 || next > len(str) {
		next = len(str)
	}

	chunk := str[i:next]

	return chunk, next
}

func diceCountHand(hand string, targetDice map[string]int) bool {
	handDice := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	i := 0
	for i < len(hand) {
		chunk, next := findChunk(hand, i)
		c := strings.Split(chunk, " ")
		n, _ := strconv.Atoi(c[0])
		if len(c) > 1 {
			handDice[c[1]] += n
		}
		i = next + 2
	}
	if targetDice["red"] >= handDice["red"] && targetDice["green"] >= handDice["green"] && targetDice["blue"] >= handDice["blue"] {
		return true
	} else {
		return false
	}
}

func diceMinCount(hand string, minDice map[string]int) map[string]int {
	i := 0
	for i < len(hand) {
		chunk, next := findChunk(hand, i)

		c := strings.Split(chunk, " ")
		n, _ := strconv.Atoi(c[0])
		if len(c) > 1 {
			if minDice[c[1]] < n {
				minDice[c[1]] = n
			}
		}
		i = next + 2
	}
	return minDice
}

func sumSlice(n []int) int {
	result := 0
	for i := 0; i < len(n); i++ {
		result += n[i]
	}
	return result
}

func DiceGame(filePath string) {
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
	power := 0
	for fScanner.Scan() {
		game := (fScanner.Text()[strings.Index(fScanner.Text(), ":")+2:])
		sets := strings.Split(game, "; ")
		setsTested := true
		minDice := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for i := 0; i < len(sets); i++ {
			if !diceCountHand(sets[i], targetDice) {
				setsTested = false
			}
			minDice = diceMinCount(sets[i], minDice)
		}

		if setsTested {
			ids = append(ids, n)
		}
		power += minDice["red"] * minDice["green"] * minDice["blue"]
		n++
	}
	log.Println("Sum of IDs:", sumSlice(ids))
	log.Println("Dice Power:", power)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Enter filename of test file")
	}

	filePath := os.Args[1]
	DiceGame(filePath)
}

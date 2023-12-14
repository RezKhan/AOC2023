package main

import (
	"cmp"
	"fmt"
	"log"
	rf "readfile"
	"slices"
	"strconv"
	"strings"
)

var cardfaces = map[string]int{
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7,
	"8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
}

type hand struct {
	cards    []string
	bid      int
	handtype int // based on sequence, 7 to 1
	score    int
}

// hand types:
// [ 5 ]			5 of a kind: 5 of one
// [ 4 1 ]			4 of a kind: 4 of one
// [ 3 2 ]			full house: 3 of one, 2 of another
// [ 3 1 1 ]		3 of a kind: 3 of one, 1 of one, 1 of one
// [ 2 2 1 ]		2 pair: 2 of one, 2 of one, 1 of one
// [ 2 1 1 1 ]		1 pair: 2 of one, 1 of one, 1of one, 1 of one
// [ 1 1 1 1 1 ]	high card: only ones
// 					-> then check higher value order

func checkHandType(hand hand) int {
	handTypeMap := map[string]int{
		"5":     7, // 5 of a kind
		"41":    6, // 4 of a kind
		"32":    5, // full house
		"311":   4, // 3 of a kind
		"221":   3, // two pair
		"2111":  2, // pair
		"11111": 1, // high card
	}

	cards := make([]string, len(hand.cards))
	copy(cards, hand.cards)
	slices.Sort(cards)
	var sequences []int
	s := 1
	sequences = append(sequences, s)

	for i := 1; i < len(cards); i++ {
		if cards[i] == cards[i-1] {
			s++
			sequences[len(sequences)-1] = s
		} else {
			s = 1
			sequences = append(sequences, s)
		}
	}

	slices.Sort(sequences)
	slices.Reverse(sequences)
	strseq := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sequences)), ""), "[]")
	return handTypeMap[strseq]
}

func setHands(lines []string) []hand {
	var hands []hand

	for _, line := range lines {
		var h hand
		tmp := strings.Split(line, " ")
		cards := strings.Split(tmp[0], "")
		bid, err := strconv.Atoi(tmp[1])
		if err != nil {
			log.Println(err)
		}
		h.cards = cards
		h.bid = bid
		h.handtype = checkHandType(h)
		hands = append(hands, h)
	}
	return hands
}

func partOne(lines []string) {
	hands := setHands(lines)
	hands2 := make([]hand, len(hands))
	copy(hands2, hands)
	slices.SortFunc(hands2, func(a hand, b hand) int {
		return cmp.Compare(a.handtype, b.handtype)
	})
	log.Println(hands)
	log.Println(hands2)
}

func main() {
	filePath := "d07test.txt"

	lines := rf.ReadFile(filePath)
	partOne(lines)
}

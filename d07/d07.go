package main

import (
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
	handtype int // based on sequence, 7
	score    int
}

// hand types:
// [ 5 ]			5 of a kind: 5 of one
// [ 4 1 ]			4 of a kind: 4 of one -> then check higher value order
// [ 3 2 ]			full house: 3 of one, 2 of another -> then check higher value order
// [ 3 1 1 ]		3 of a kind: 3 of one, 1 of one, 1 of one -> then check higher value order
// [ 2 2 1 ]		2 pair: 2 of one, 2 of one, 1 of one -> then check higher value order
// [ 2 1 1 1 ]		1 pair: 2 of one, 1 of one, 1of one, 1 of one  -> then check higher value order
// [ 1 1 1 1 1 ]	Only ones -> then check higher value order

func checkHandType(hand hand) {
	cards := make([]string, len(hand.cards))
	copy(cards, hand.cards)
	slices.Sort(cards)
	log.Println(hand.cards, cards)
	var sequences []int
	s := 1
	for i := 0; i < len(cards); i++ {
		if i == 0 {
			sequences = append(sequences, s)
			continue
		}
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
	log.Println(sequences)
}

func partOne(lines []string) {
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
		checkHandType(h)
		hands = append(hands, h)
	}
}

func main() {
	filePath := "d07test.txt"

	lines := rf.ReadFile(filePath)
	partOne(lines)
}

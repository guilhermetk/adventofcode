package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	task1()
}

var cardOrder map[string]int = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

var hands []Hand

//five of a kind:	AAAAA
//four of a kind:	AAAAT
//full house:		AAATT
//three of a kind:	AAATY
//two pair:			AATTX
//one pair:			AATXY
//high card:		23456

type Hand struct {
	cards [5]string
	bid   int
}

func task1() {
	// fileContent, _ := os.ReadFile("./edge.txt")
	fileContent, _ := os.ReadFile("./test.txt")
	fileContent, _ := os.ReadFile("./prod.txt")
	lines := strings.Split(string(fileContent), "\n")
	for _, line := range lines {
		if len(line) > 0 {
			cardSplit := strings.Split(line, " ")
			cards := strings.Split(cardSplit[0], "")
			bid, _ := strconv.Atoi(cardSplit[1])
			hand := Hand{cards: [5]string(cards), bid: bid}
			hands = append(hands, hand)
		}
	}
	slices.SortFunc(hands, sortByKind)
	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += (i + 1) * hands[i].bid
		fmt.Println(i+1, hands[i].cards, hands[i].bid, hands[i].kind())
	}
	fmt.Println(sum)
}

func sortByKind(a, b Hand) int {
	if a.kind() < b.kind() {
		return -1
	}
	if a.kind() > b.kind() {
		return 1
	}

	return sortHandByCards(a, b)
}

func sortHandByCards(a, b Hand) int {
	for i := 0; i < 5; i++ {
		aCard := cardOrder[a.cards[i]]
		bCard := cardOrder[b.cards[i]]
		if aCard > bCard {
			return 1
		}
		if aCard < bCard {
			return -1
		}
	}
	return 0
}

func (h Hand) kind() int {
	groups := map[string]int{}
	jokers := 0
	for _, card := range h.cards {
		if card != "J" {
			if groups[card] == 0 {
				groups[card] = 1
			} else {
				groups[card] += 1
			}

		} else {
			jokers++
		}
	}
	fmt.Println(groups)
	if jokers > 0 {
		delete(groups, "J")
		highestGroupKey := ""
		highestGroupCount := 0
		for key, value := range groups {
			if value >= highestGroupCount {
				highestGroupCount = value
				highestGroupKey = key
			}
		}
		groups[highestGroupKey] += jokers
	}
	fmt.Println(groups)
	switch len(groups) {
	case 5:
		return 1 //high card
	case 1:
		return 7 // five of a kind
	case 2:
		var val int
		for _, v := range groups {
			val = v
			break
		}
		if val == 4 || val == 1 {
			return 6 // four of a kind
		} else {
			return 5 //full house
		}
	case 3:
		for _, value := range groups {
			if value == 2 {
				return 3 //two pair
			} else if value == 3 {
				return 4 //three pair
			}
		}
	case 4:
		return 2 //one pair
	}
	return 0
}

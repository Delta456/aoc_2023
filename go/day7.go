package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Hand struct {
	hand            string
	handType        int
	morphedHand     string
	morphedHandType int
	bid             int
}

func main() {
	buff, err := os.ReadFile("../input/day7_sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := strings.TrimSpace(string(buff[:]))
	fileArr := strings.Split(fileContent, "\n")

	var handStr string
	var bid int
	var processedHands []Hand

	for _, content := range fileArr {
		fmt.Sscanf(content, "%s %d", &handStr, &bid)
		processedHands = append(processedHands, Hand{
			hand:     handStr,
			handType: getHandType(handStr),
			bid:      bid,
		})

	}
	slices.SortFunc(processedHands, cmpHands)
	totalWins := 0
	for i, hand := range processedHands {
		totalWins += (i + 1) * hand.bid
	}
	fmt.Println(totalWins)
}

func getUniqueCardsFromHand(hand string) []string {
	var uniqueCards []string
	for _, card := range hand {
		if !slices.Contains(uniqueCards, string(card)) {
			uniqueCards = append(uniqueCards, string(card))
		}
	}
	return uniqueCards
}

func getHandType(hand string) int {
	handType := map[string]int{"HighCard": 1, "OnePair": 2, "TwoPair": 3, "ThreeofKind": 4, "FullHouse": 5, "FourofKind": 6, "FiveofKind": 7}

	uniqueCards := getUniqueCardsFromHand(hand)

	if len(uniqueCards) == 1 {
		return handType["FiveofKind"]
	}

	if len(uniqueCards) == 2 {
		for _, card := range uniqueCards {
			countHands := strings.Count(hand, card)
			if countHands == 4 {
				return handType["FourofKind"]
			} else if countHands == 3 {
				return handType["FullHouse"]
			}
		}
	}

	if len(uniqueCards) == 3 {
		for _, card := range uniqueCards {
			countHands := strings.Count(hand, card)
			if countHands == 3 {
				return handType["ThreeofKind"]
			} else if countHands == 2 {
				return handType["TwoPair"]
			}

		}
	}
	if len(uniqueCards) == 4 {
		for _, card := range uniqueCards {
			if strings.Count(hand, card) == 2 {
				return handType["OnePair"]
			}
		}
	}

	if len(uniqueCards) == 5 {
		return handType["HighCard"]
	}

	return -1
}

func cmpHands(hand1, hand2 Hand) int {
	var result int
	var cardPriority map[string]int

	cardPriority = map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	if hand1.handType == hand2.handType {
		for i := 0; i < len(hand1.hand); i++ {
			if cardPriority[string(hand1.hand[i])] > cardPriority[string(hand2.hand[i])] {
				result = 1
				break
			} else if cardPriority[string(hand1.hand[i])] < cardPriority[string(hand2.hand[i])] {
				result = -1
				break
			}
		}
	} else if hand1.handType > hand2.handType {
		result = 1
	} else if hand1.handType < hand2.handType {
		result = -1
	}

	return result
}

func cmpHandsWithNewRules(h1, h2 Hand) int {
	var result int

	cardPriority := map[string]int{
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

	if h1.morphedHandType == h2.morphedHandType {
		for i := 0; i < len(h1.hand); i++ {
			if cardPriority[string(h1.hand[i])] > cardPriority[string(h2.hand[i])] {
				result = 1
				break
			} else if cardPriority[string(h1.hand[i])] < cardPriority[string(h2.hand[i])] {
				result = -1
				break
			}
		}
	} else if h1.morphedHandType > h2.morphedHandType {
		result = 1
	} else if h1.morphedHandType < h2.morphedHandType {
		result = -1
	}

	return result
}

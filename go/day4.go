package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	buff, err := os.ReadFile("../input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(buff[:])
	totalSum := 0
	var cardWinsNoOfCards = make(map[int]int)
	var cardQueue []int

	for i, content := range strings.Split(fileContent, "\n") {

		colorCards := strings.Split(content, ": ")[1]
		deck := strings.Split(colorCards, " | ")
		winningCards := strings.Split(deck[0], " ")
		cards := strings.Split(deck[1], " ")

		actualCards := intersection(winningCards, cards)

		sum := 1
		if len(actualCards) >= 1 {
			sum += int(math.Pow(2, float64(len(actualCards)-1))) - 1
		} else if len(actualCards) == 0 {
			sum = 0
		}
		cardWinsNoOfCards[i+1] = len(actualCards)
		cardQueue = append(cardQueue, i+1)
		totalSum += sum

	}
	var cardsProcessed int
	var queueLen = len(cardQueue)

	for cardsProcessed < queueLen {
		start := cardQueue[cardsProcessed]
		end := cardQueue[cardsProcessed] + cardWinsNoOfCards[cardQueue[cardsProcessed]]
		var genRange = make([]int, 0, end-start+1)
		for i := start + 1; i <= end; i++ {
			genRange = append(genRange, i)

		}
		cardQueue = append(cardQueue, genRange...)
		cardsProcessed++
		queueLen += len(genRange)
	}
	fmt.Println("Part 1:", totalSum)
	fmt.Println("Part 2:", cardsProcessed)

}

func intersection(winningCards, cards []string) []string {
	var common []string
	for _, ele := range cards {
		for _, subEle := range winningCards {
			if ele == subEle && subEle != "" {
				common = append(common, subEle)
			}
		}
	}
	return common
}

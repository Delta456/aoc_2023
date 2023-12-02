package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	ID    string
	Red   int
	Blue  int
	Green int
}

func (g Game) String() string {
	return fmt.Sprintf("Game %s: Red ball: %d Blue Ball: %d Green ball: %d", g.ID, g.Red, g.Blue, g.Green)
}

func main() {
	buff, err := os.ReadFile("../input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(buff[:])
	var result int
	var games []Game

	for _, content := range strings.Split(fileContent, "\n") {

		before, after, found := strings.Cut(content, ":")
		if found {
			ID := before[5:]
			game := Game{ID: ID}

			removeNewLines := strings.Split(after, "\n")
			var totalRed, totalBlue, totalGreen []int
			for _, str := range removeNewLines {
				var index int
				for index < len(str) {
					stringAlpha := string(str[index])
					if isDigit(stringAlpha) {
						if isDigit(string(str[index+1])) {
							if str[index+3] == 'b' {
								blue, err := strconv.Atoi(stringAlpha + string(str[index+1]))
								if err != nil {
									log.Fatal(err)
								}
								totalBlue = append(totalBlue, blue)
								index += 2
							}
						} else {
							if str[index+2] == 'b' {
								blue, err := strconv.Atoi(stringAlpha)
								if err != nil {
									log.Fatal(err)
								}
								totalBlue = append(totalBlue, blue)
							}
						}
						if isDigit(string(str[index+1])) {
							if str[index+3] == 'r' {
								red, err := strconv.Atoi(stringAlpha + string(str[index+1]))
								if err != nil {
									log.Fatal(err)
								}
								totalRed = append(totalRed, red)
								index += 2
							}
						} else {
							if str[index+2] == 'r' {
								red, err := strconv.Atoi(stringAlpha)
								if err != nil {
									log.Fatal(err)
								}
								totalRed = append(totalRed, red)
							}
						}
						if isDigit(string(str[index+1])) {
							if str[index+3] == 'g' {
								green, err := strconv.Atoi(stringAlpha + string(str[index+1]))
								if err != nil {
									log.Fatal(err)
								}
								totalGreen = append(totalGreen, green)
								index += 2
							}
						} else {
							if str[index+2] == 'g' {
								green, err := strconv.Atoi(stringAlpha)
								if err != nil {
									log.Fatal(err)
								}
								totalGreen = append(totalGreen, green)
							}
						}
					}
					index++

				}
				game.Blue = slices.Max(totalBlue)
				game.Red = slices.Max(totalRed)
				game.Green = slices.Max(totalGreen)
				result += game.Red * game.Blue * game.Green
				games = append(games, game)
			}

		}
	}
	var sum int
	for _, g := range games {
		if g.Red <= 12 && g.Green <= 13 && g.Blue <= 14 {
			id, err := strconv.Atoi(g.ID)
			if err != nil {
				log.Fatal(err)
			}
			sum += id
		}
	}
	fmt.Println(sum, result)
}

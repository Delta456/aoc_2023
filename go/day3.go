package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	buff, err := os.ReadFile("../input/day3_sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(buff[:])
	cells := make([][]rune, len(strings.Split(fileContent, "\n")))

	for i := 0; i < len(cells); i++ {
		cells[i] = make([]rune, len(strings.Split(fileContent, "\n")[0]))
	}

	for i, content := range strings.Split(fileContent, "\n") {
		for j, str := range content {
			cells[i][j] = str
		}
	}

	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells); y++ {
			if !unicode.IsNumber(cells[x][y]) && cells[x][y] != '.' {

			}
		}
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	buff, err := os.ReadFile("../input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(buff[:])
	var sum int

	for _, content := range strings.Split(fileContent, "\n") {
		count := 0
		var s string

		for index, str := range content {
			contentString := string(str)

			if isDigit(contentString) {
				count++
				s += contentString
			} else {
				for k := range numbers {
					if len(content[index:]) >= len(k) && content[index:index+len(k)] == k {
						s += numbers[k]
						count++
					}
				}
			}

		}
		if count > 2 {
			a := string(s[0]) + string(s[len(s)-1])
			num, err := strconv.Atoi(a)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
		} else if count != 0 {
			if count == 1 {
				s += s
			}
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
		}
	}
	fmt.Println(sum)
}

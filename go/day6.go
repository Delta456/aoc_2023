package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	buff, err := os.ReadFile("../input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(buff[:])
	details := strings.Split(fileContent, "\n")
	time := convertArrToIntArr(strings.Fields(details[0])[1:])
	dist := convertArrToIntArr(strings.Fields(details[1])[1:])

	timeField, _ := strconv.Atoi(strings.Join(strings.Fields(details[0])[1:], ""))
	distField, _ := strconv.Atoi(strings.Join(strings.Fields(details[1])[1:], ""))
	//fmt.Println(timeField, distField)
	product := 1

	var index int
	for index < len(time) {
		waysToWin := 0
		for timeToPushButton := 1; timeToPushButton < time[index]; timeToPushButton++ {
			remainTime := time[index] - timeToPushButton
			//fmt.Println(fmt.Sprintf("%d := %d - %d", remainTime, time[index], timeToPushButton))
			distance := timeToPushButton * remainTime
			//fmt.Println(fmt.Sprintf("%d := %d * %d\n", distance, timeToPushButton, remainTime))
			if distance > dist[index] {
				waysToWin++
			}
		}
		//fmt.Println(waysToWin)
		product *= waysToWin
		index++
	}

	// a = -1 due to the step time being constant and a < 0 indicates a
	// parabola open downward (upside down U)
	a := float64(-1)
	// b is the duration of the race, and we shift the parabola to the right of
	// the y-axis
	b := float64(timeField)
	// c shifts the roots of the parabola to the edges of the record we need to
	// beat
	c := float64(-(distField + 1))

	// Use the quadratic formula to find the x roots
	x1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	x2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	ans := int(math.Floor(x2) - math.Ceil(x1) + 1)
	fmt.Println("Part 1: ", product)
	fmt.Println("Part 2: ", ans)

}

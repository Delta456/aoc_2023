package main

import (
	"strconv"
)

func isDigit(digit string) bool {
	_, err := strconv.Atoi(digit)
	return err == nil
}

var numbers = map[string]string{"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func convertArrToIntArr(strArr []string) []int {
	var intArr []int
	for _, str := range strArr {
		i, _ := strconv.Atoi(str)

		intArr = append(intArr, i)
	}
	return intArr
}

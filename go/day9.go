package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	buff, err := os.ReadFile("../input/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := strings.Split(string(buff[:]), "\n")

	input := make([][]int, 0)

	for _, content := range fileContent {
		var nums []int
		arr := strings.Fields(content)
		for _, ele := range arr {
			num, _ := strconv.Atoi(ele)
			nums = append(nums, num)
		}
		input = append(input, nums)
	}

	fmt.Println("Part 1: ", extrapolate(input))

	// 2D arrays cannot be reversed with just slices.Reverse(input)
	// Need to reverse nested arrays via loop
	for _, arr := range input {
		slices.Reverse(arr)
	}

	fmt.Println("Part 2: ", extrapolate(input))
}

func extrapolate(input [][]int) (sum int) {
	for _, arr := range input {
		var stop bool

		for !stop {
			stop = true
			sum += arr[len(arr)-1]
			newArr := make([]int, len(arr)-1)

			for i := 0; i < len(arr)-1; i++ {
				newArr[i] = arr[i+1] - arr[i]
				if !slicesHasAllZeros(newArr) {
					stop = false
				}
			}
			arr = newArr
		}
	}
	return sum
}

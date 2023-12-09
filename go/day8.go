package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Navigate struct {
	Name  string
	Left  string
	Right string
}

func main() {
	buff, err := os.ReadFile("../input/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := strings.Split(string(buff[:]), "\n")

	direction := fileContent[0]
	var navigator = make(map[string]Navigate)

	for _, content := range fileContent[2:] {
		var navi Navigate
		_, err := fmt.Sscanf(content, "%s = (%3s, %3s)", &navi.Name, &navi.Left, &navi.Right)
		if err != nil {
			log.Fatal(err)
		}
		navigator[navi.Name] = navi
	}

	current := navigator["AAA"]
	fmt.Println(traverse(direction, current, navigator))

}

func traverse(direction string, curr Navigate, navi map[string]Navigate) int {
	steps := 0

	for {
		if strings.HasSuffix(curr.Name, "Z") {
			return steps
		}
		switch direction[steps%len(direction)] {
		case 'R':
			curr = navi[curr.Right]
		case 'L':
			curr = navi[curr.Left]
		}
		steps++
	}
}

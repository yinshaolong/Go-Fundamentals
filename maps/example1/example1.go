// Sample program to show how maps behave when you read an
// absent key.
package main

import (
	"fmt"
	"strings"
)

func main() {
	repeatText := RepeatText("*", 20)
	players := make(map[string]int)
	score := players["anna"]
	fmt.Println("Score:", score, "Current map:", players)
	fmt.Println(repeatText)
	players["anna"]++
	score, ok := players["anna"]
	fmt.Println("Score:", score, "Present: ", ok, "Current map:", players)
	fmt.Println(repeatText)

	//traditional method of determining if the key is absent:
	if n, ok := players["anna"]; ok {
		players["anna"] = n + 1
	} else {
		players["anna"] = n
	}
}

func RepeatText(text string, number int) string {
	line := make([]string, number)
	for n := 0; n < number; n++ {
		line[n] = text
	}
	return strings.Join(line, "")
}

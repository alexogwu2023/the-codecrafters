package main

import (
	"fmt"
	"strings"
)

func LowerCase(text string) string {
	words := strings.Fields(text)
	for i, s := range words {
		if s == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}

		//words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, " ")

}
func main() {
	fmt.Println(LowerCase("hope IT WAS THE FIRST (low, 2) come"))
}

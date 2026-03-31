package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UpperText(word string) string {
	return strings.ToUpper(word)
}

func LowerText(s string) string {
	return strings.ToLower(s)
}

func UpperFirstAlphabet(s string) string {
	word := strings.ToLower(s)
	newword := strings.Title(word)
	return newword
}

func ReverseWord(text string) string {
	words := strings.Split(text, " ")
	for i, word := range words {
		runes := []rune(word)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func SnakeCase(words string) string {
	words = strings.ReplaceAll(words, "!", " ")
	words = strings.ReplaceAll(words, " ", "_")
	return strings.ToLower(strings.ReplaceAll(words, ",", "_"))
}

func TitleCase(text string) string {
	smallwords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true, "or": true,
		"for": true, "nor": true, "on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true, "is": true, "it": true,
	}

	words := strings.Fields(text)
	for i, word := range words {
		lower := strings.ToLower(word)
		if i == 0 || !smallwords[lower] {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		} else {
			words[i] = lower
		}
	}

	return strings.Join(words, " ")

}
func main() {
	for {
		fmt.Println("Enter a word: ")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')

		var editor string
		fmt.Println("Choose Operation(upper, lower, cap, title, snake, reverse, exit): ")
		fmt.Scanln(&editor)

		if editor == "exit" {
			fmt.Println("exiting, goodbye")
			break
		}

		if editor == "upper" {
			fmt.Println(UpperText(name))
		}
		if editor == "lower" {
			fmt.Println(LowerText(name))
		}
		if editor == "cap" {
			fmt.Println(UpperFirstAlphabet(name))
		}
		if editor == "title" {
			fmt.Println(TitleCase(name))
		}
		if editor == "snake" {
			fmt.Println(SnakeCase(name))
		}
		if editor == "reverse" {
			fmt.Println(ReverseWord(name))
		}

	}
}

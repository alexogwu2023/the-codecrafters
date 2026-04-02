package main

import (
	"fmt"
	"strings"
	"unicode"
	"os"
)

func trimWhitespace(text string) string {
	return strings.TrimSpace(text)
}

func replaceTODO(text string) string {
	return strings.ReplaceAll(text, "TODO:", "ACTION:")
}

func allCapsToTitleCase(text string) string {
	trimmed := strings.TrimSpace(text)

	// Check if line is fully uppercase
	if trimmed == strings.ToUpper(trimmed) && trimmed != "" {
		words := strings.Fields(strings.ToLower(trimmed))

		for i := 0; i < len(words); i++ {
			runes := []rune(words[i])
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
				words[i] = string(runes)
			}
		}

		return strings.Join(words, " ")
	}

	return text
}

func lowercaseToUppercase(text string) string {
	trimmed := strings.TrimSpace(text)

	// Check if line is fully lowercase
	if trimmed == strings.ToLower(trimmed) && trimmed != "" {
		return strings.ToUpper(text)
	}

	return text
}

func reverseWordsIfContainsReverse(text string) string {
	if !strings.Contains(text, "REVERSE") {
		return text
	}

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		words[i] = reverseWord(words[i])
	}

	return strings.Join(words, " ")
}

func reverseWord(word string) string {
	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func printSummary(read, written, removed int) {
	fmt.Println("✦ Lines read    :", read)
	fmt.Println("✦ Lines written :", written)
	fmt.Println("✦ Lines removed :", removed)
	fmt.Println("✦ Rules applied : trimWhitespace, replaceTODO, titleCase, uppercase, reverseWords")
}



func main() {
	// Rule 2: Check arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("✗ File not found:", inputFile)
		return
	}

	content := string(data)

	// Handle empty file safely
	if strings.TrimSpace(content) == "" {
		err = os.WriteFile(outputFile, []byte(""), 0644)
		if err != nil {
			fmt.Println("✗ Cannot write output file:", err)
			return
		}
		printSummary(0, 0, 0)
		return
	}

	lines := strings.Split(content, "\n")

	var outputLines []string
	linesRead := len(lines)

	// Rule 3: Apply 5 transformation rules
	for _, line := range lines {
		line = trimWhitespace(line)
		line = replaceTODO(line)
		line = allCapsToTitleCase(line)
		line = lowercaseToUppercase(line)
		line = reverseWordsIfContainsReverse(line)

		outputLines = append(outputLines, line)
	}

	result := strings.Join(outputLines, "\n")

	// Rule 4: Write output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("✗ Cannot write output file:", err)
		return
	}

	linesWritten := len(outputLines)
	linesRemoved := linesRead - linesWritten

	printSummary(linesRead, linesWritten, linesRemoved)
}
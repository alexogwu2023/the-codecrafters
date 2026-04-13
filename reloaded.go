/*package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ProcessText runs all the text modifications required by the project.
func ProcessText(text string) string {

	// Convert hexadecimal numbers before (hex)
	text = processHex(text)

	// Convert binary numbers before (bin)
	text = processBin(text)

	// Apply case modifications: (up), (low), (cap)
	text = processCase(text)

	// Fix punctuation spacing rules
	text = fixPunctuation(text)

	// Fix quotation marks spacing
	text = fixQuotes(text)

	// Convert "a" to "an" when needed
	text = fixArticles(text)

	return text
}

// HEX PROCESSING

// processHex converts hexadecimal numbers followed by "(hex)"
// into their decimal representation.
func processHex(text string) string {

	// Regex to match: hexadecimal number followed by "(hex)"
	re := regexp.MustCompile(`(\b[0-9A-Fa-f]+\b) \(hex\)`)

	// Replace matches using a function
	return re.ReplaceAllStringFunc(text, func(match string) string {

		// Split the match into parts
		parts := strings.Split(match, " ")

		// Convert hexadecimal to decimal
		val, _ := strconv.ParseInt(parts[0], 16, 64)

		// Return the decimal value
		return fmt.Sprintf("%d", val)
	})
}

// BINARY PROCESSING

// processBin converts binary numbers followed by "(bin)"
// into their decimal representation.
func processBin(text string) string {

	// Regex to match binary number followed by "(bin)"
	re := regexp.MustCompile(`(\b[01]+\b) \(bin\)`)

	return re.ReplaceAllStringFunc(text, func(match string) string {

		// Split match into parts
		parts := strings.Split(match, " ")

		// Convert binary to decimal
		val, _ := strconv.ParseInt(parts[0], 2, 64)

		return fmt.Sprintf("%d", val)
	})
}

// CASE MODIFICATIONS

// processCase handles the following markers:
// (up)  -> uppercase previous word
// (low) -> lowercase previous word
// (cap) -> capitalize previous word
func processCase(text string) string {

	// Split the text into words
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		switch words[i] {

		case "(up)":
			// Convert previous word to uppercase
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}

			// Remove the marker from the slice
			words = remove(words, i)
			i--

		case "(low)":
			// Convert previous word to lowercase
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}

			words = remove(words, i)
			i--

		case "(cap)":
			// Capitalize the previous word
			if i > 0 {
				words[i-1] = strings.Title(words[i-1])
			}

			words = remove(words, i)
			i--
		}
	}

	// Rejoin the words into a sentence
	return strings.Join(words, " ")
}

// remove removes an element from a slice at index i.
func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

// PUNCTUATION FORMATTING

// fixPunctuation ensures punctuation marks follow formatting rules:
// punctuation attaches to the previous word and has a space after it.
func fixPunctuation(text string) string {

	// Remove spaces before punctuation
	re := regexp.MustCompile(`\s+([.,!?;:])`)
	text = re.ReplaceAllString(text, "$1")

	// Ensure there is a space after punctuation
	re2 := regexp.MustCompile(`([.,!?;:])([^\s])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	return text
}

// QUOTES FIXING

// fixQuotes removes spaces inside single quotes.
func fixQuotes(text string) string {

	// Match text inside quotes and remove surrounding spaces
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)

	return re.ReplaceAllString(text, "'$1'")
}

// ARTICLE CORRECTION

// fixArticles converts "a" to "an" when the next word
// begins with a vowel or 'h'.
func fixArticles(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {

		// Check if the current word is "a"
		if strings.ToLower(words[i]) == "a" {

			next := strings.ToLower(words[i+1])

			// Check if next word begins with vowel or 'h'
			if strings.HasPrefix(next, "a") ||
				strings.HasPrefix(next, "e") ||
				strings.HasPrefix(next, "i") ||
				strings.HasPrefix(next, "o") ||
				strings.HasPrefix(next, "u") ||
				strings.HasPrefix(next, "h") {

				// Preserve capitalization
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}*/

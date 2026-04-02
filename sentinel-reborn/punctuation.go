package main

import (
	"strings"
)

// DONE
func multiPunctuation(text string) string {
	replacements := map[string]string{
		" ...": "...",
		" !!":  "!!",
		" ??":  "??",
		" !?":  "!?",
		" ?!":  "?!",
	}

	for old, newVal := range replacements {
		text = strings.ReplaceAll(text, old, newVal)
	}

	return text
}

// DONE
func fixPunctuation(text string) string {
	text = multiPunctuation(text)

	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {
		text = strings.ReplaceAll(text, " "+p, p+" ")
		text = strings.ReplaceAll(text, p+" ", p+" ")
	}

	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")

	return strings.Join(strings.Fields(text), " ")
}

// DONE
func fixArticle(word string) string {
 words := strings.Fields(word)
 vowels := "aeiouhAEIOUH"
 for i := 0; i < len(words)-1; i++ {
     if (words[i] == "a" || words[i] == "A") && strings.ContainsRune(vowels, rune(words[i+1][0])) {
         if words[i] == "A" {
             words[i] = "An"
         } else {
             words[i] = "an"
         }
     }
 }
 return strings.Join(words, " ")
}
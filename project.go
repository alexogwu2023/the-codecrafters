package main

import (
	//"fmt"
	"strconv"
	"strings"
)

func processText(text string) string {
	text = convertHexBin(text)

	text = caseProcessing(text)

	text = FixPunctuation(text)

	text = FixQuotes(text)

	text = FixArticles(text)

	return text
}

func convertHexBin(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			n, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(bin)" && i > 0 {
			b, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(b, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func caseProcessing(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(up," {
			words[i+1] = strings.TrimSuffix(words[i+1], ")")
			nwords, err := strconv.Atoi(words[i+1])
			if err == nil {
				for n := i - 1; n >= i-nwords && n >= 0; n-- {
					words[n] = strings.ToUpper(words[n])
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}

		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(low," {
			words[i+1] = strings.TrimSuffix(words[i+1], ")")
			nwords, err := strconv.Atoi(words[i+1])
			if err == nil {
				for n := i - 1; n >= i-nwords && n >= 0; n-- {
					words[n] = strings.ToLower(words[n])
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}

		if words[i] == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(cap," {
			words[i+1] = strings.TrimSuffix(words[i+1], ")")
			nwords, err := strconv.Atoi(words[i+1])
			if err == nil {
				for n := i - 1; n >= i-nwords && n >= 0; n-- {
					words[n] = strings.Title(words[n])
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}
	}
	return strings.Join(words, " ")

}

func FixArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if words[i] == "a" || words[i] == "A" {
			if i+1 < len(words) && strings.ContainsRune("aeiouhAEIOUH", rune(words[i+1][0])) {
				if words[i] == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}

		}
	}
	return strings.Join(words, " ")
}

func IsPunctuation(t string) bool {
	for _, i := range t {
		if !strings.ContainsRune(",.;:'!?", rune(i)) {
			return false
		}
	}
	return true
}

func FixPunctuation(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if IsPunctuation(words[i]) {
			words[i-1] = words[i-1] + words[i]
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		for len(words[i]) > 0 && IsPunctuation(string(words[i][0])) {
			words[i-1] += string(words[i][0])
			words[i] = words[i][1:]
		}

	}

	return strings.Join(words, " ")
}

func FixQuotes(text string) string {
	for {
		first := strings.Index(text, "'")
		if first == -1 {
			break
		}
		second := strings.Index(text[first+1:], "'")
		if second == -1 {
			break
		}

		second = second + first + 1

		before := text[:first]
		middle := text[first+1 : second]
		after := text[second+1:]

		trimmedMiddle := strings.TrimSpace(middle)

		text = before + "'" + trimmedMiddle + "'" + after
		break
	}
	return text
}

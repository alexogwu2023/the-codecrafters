package main

import (
	"strconv"
	"strings"
	"unicode"
)


// DONE
func converter(input string) string {
	words := strings.Fields(input)

	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {
			if n, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--

		} else if words[i] == "(bin)" {
			if n, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

func capitalizeWord(word string) string {
    runes := []rune(strings.ToLower(word))
    if len(runes) > 0 {
        runes[0] = unicode.ToUpper(runes[0])
    }
    return string(runes)
}

func toLowerCase(text string) string {
    words := strings.Fields(text)

    for i := 0; i < len(words); i++ {
        if words[i] == "(low)" {
            if i-1 >= 0 {
                words[i-1] = strings.ToLower(words[i-1])
            }
            words = append(words[:i], words[i+1:]...)
            i--
            continue
        }

        if words[i] == "(low," && i+1 < len(words) {
            countStr := strings.TrimSuffix(words[i+1], ")")
            n, err := strconv.Atoi(countStr)
            if err == nil {
                start := i - n
                if start < 0 {
                    start = 0
                }
                for j := start; j < i; j++ {
                    words[j] = strings.ToLower(words[j])
                }
            }
            words = append(words[:i], words[i+2:]...)
            i--
        }
    }

    return strings.Join(words, " ")
}

func toUpperCase(text string) string {
    words := strings.Fields(text)

    for i := 0; i < len(words); i++ {
        if words[i] == "(up)" {
            if i-1 >= 0 {
                words[i-1] = strings.ToUpper(words[i-1])
            }
            words = append(words[:i], words[i+1:]...)
            i--
            continue
        }

        if words[i] == "(up," && i+1 < len(words) {
            countStr := strings.TrimSuffix(words[i+1], ")")
            n, err := strconv.Atoi(countStr)
            if err == nil {
                start := i - n
                if start < 0 {
                    start = 0
                }
                for j := start; j < i; j++ {
                    words[j] = strings.ToUpper(words[j])
                }
            }
            words = append(words[:i], words[i+2:]...)
            i--
        }
    }

    return strings.Join(words, " ")
}

func toTitleCase(text string) string {
    words := strings.Fields(text)

    for i := 0; i < len(words); i++ {
        if words[i] == "(cap)" {
            if i-1 >= 0 {
                words[i-1] = capitalizeWord(words[i-1])
            }
            words = append(words[:i], words[i+1:]...)
            i--
            continue
        }

        if words[i] == "(cap," && i+1 < len(words) {
            countStr := strings.TrimSuffix(words[i+1], ")")
            num, err := strconv.Atoi(countStr)
            if err == nil {
                start := i - num
                if start < 0 {
                    start = 0
                }
                for j := start; j < i; j++ {
                    words[j] = capitalizeWord(words[j])
                }
            }
            words = append(words[:i], words[i+2:]...)
            i--
        }
    }

    return strings.Join(words, " ")
}

func applyTransformation(s string) string {
	s = converter(s)
	s = fixArticle(s)
	s = fixPunctuation(s)
	s = toUpperCase(s)
	s = toLowerCase(s)
	s = toTitleCase(s)

	return s
}
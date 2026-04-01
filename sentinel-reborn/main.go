package main

import (
	"fmt"
	"os"
	"strings"
)

func LowerCase(text string) string {
	words := strings.Fields(text)
	for i, s := range words {
		if s == "(low," {
			n := int(words[i+1][0] - '0')

			for j := i - n; j < i; j++ {
				if j >= 0 {
					words[j] = strings.ToLower(words[j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			break
			
		}

	}
	return strings.Join(words, " ")

}
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go input.txt output.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	result := LowerCase(string(data))

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
		return
	}
}
package main

import (
	"fmt"
	"os"
	//"strings"
)

/*
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
*/
/*func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading sample file:", err)
		return
	}

	result := applyTransformation(string(data))

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing result file:", err)
		return
	}
}*/

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	return string(data)
}

func writeFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
	}

	input1 := os.Args[1]
	input2 := os.Args[2]

	content := readFile(input1)
	content = applyTransformation(content)
	writeFile(input2, content)
}

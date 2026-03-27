package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var input, base string

		fmt.Print("choose operation: ")
		fmt.Print("hex: ")
		fmt.Print("bin: ")
		fmt.Print("dec: ")

		fmt.Scan(&input)

		if input == "quit" {
			break
		}

		fmt.Scan(&base)
		base = strings.ToLower(base)

		if base == "hex" {
			val, err := strconv.ParseInt(input, 16, 64)
			if err != nil {
				fmt.Println(" Invalid hex input")
				continue
			}
			fmt.Println(" Decimal:", val)

		} else if base == "bin" {
			val, err := strconv.ParseInt(input, 2, 64)
			if err != nil {
				fmt.Println(" Invalid binary input")
				continue
			}
			fmt.Println(" Decimal:", val)

		} else if base == "dec" {
			val, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("✗ Invalid decimal input")
				continue
			}
			fmt.Println("✦ Binary: ", strconv.FormatInt(val, 2))
			fmt.Println("✦ Hex:    ", strings.ToUpper(strconv.FormatInt(val, 16)))

		} else {
			fmt.Println("✗ Unknown base (use hex, bin, dec)")
		}
	}
}

/*package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var num1 string
		var base int

		fmt.Println("Base Converter: ")
		fmt.Println(" Choose Operation\n")
		fmt.Println(" 1. Convert to Hex: ")
		fmt.Println(" 2. Convert to Bin: ")
		fmt.Println(" 3. Convert to Dec: ")
		fmt.Scan(&num1)

		fmt.Print("Input Number: ")
		fmt.Scan(&base)

		var Operation string
		fmt.Print("select operation (Hex,Bin,Dec,Quit): ")
		fmt.Scanln(&Operation)

		if Operation == "Quit" {
			break
		}

		if Operation == "Hex" {
			fmt.Println("Decimal: "  strconv.ParseInt(num1, 16, 64))
		}
	}
}*/

/*if base == "hex" {
			value, err := strconv.ParseInt(num1, 16, 64)
		}
		if err != nil {
			fmt.Println("Invalid Input")
			continue
		}
		fmt.Println("Decimal: ", value)

	}
}*/

/*func ConvertToDec(number string, base int) (int64, error) {
	return strings.ParseInt(number, base 64)
}

func ConvertToBin(number, )*/

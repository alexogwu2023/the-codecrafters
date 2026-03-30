package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToDec(hex string) (int64, error) {
	return strconv.ParseInt(hex, 16, 64)

}

func BinaryToDec(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 64)
}

func DecimalToHex(dec int64) string {
	return strconv.FormatInt(dec, 16)
}

func DecimalToBin(dec int64) string {
	return strconv.FormatInt(dec, 2)
}

func main() {
	for {
		var input string
		var Converter string

		fmt.Print("Enter word: ")
		fmt.Scanln(&input)

		fmt.Print("Enter Base (hex, bin, dec, exit): ")
		fmt.Scanln(&Converter)

		if Converter == "exit" {
			fmt.Println("..exited..")
			break
		}

		if Converter != "hex" && Converter != "bin" && Converter != "dec" && Converter != "exit" {
			fmt.Println("..Invalid Operation")
			continue
		}
		if Converter == "hex" {
			val, err := HexToDec(input)
			if err != nil {
				fmt.Println("Not valid hex")
			} else {
				fmt.Println("Decimal: ", val)
			}
		}
		if Converter == "bin" {
			val, err := BinaryToDec(input)
			if err != nil {
				fmt.Println("Not valid bin")
			} else {
				fmt.Println("Decimal: ", val)
			}
		}
		if Converter == "dec" {
			val, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Not valid Decimal")
			}
			fmt.Println(strings.ToUpper(DecimalToHex(val)))

			i, _ := strconv.ParseInt(input, 10, 64)
			fmt.Println(DecimalToBin(i))

		}
	}

}

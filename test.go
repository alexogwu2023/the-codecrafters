package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var input, base string

		fmt.Print("> convert (or quit): ")
		fmt.Scan(&input)

		if input == "quit" {
			break
		}

		fmt.Scan(&base)
		base = strings.ToLower(base)

		if base == "hex" {
			val, err := strconv.ParseInt(input, 16, 64)
			if err != nil {
				fmt.Println("✗ Invalid hex input")
				continue
			}
			fmt.Println("✦ Decimal:", val)

		} else if base == "bin" {
			val, err := strconv.ParseInt(input, 2, 64)
			if err != nil {
				fmt.Println("✗ Invalid binary input")
				continue
			}
			fmt.Println("✦ Decimal:", val)

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
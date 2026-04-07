package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(" === GOPHERS CLI CALCULATOR ONLINE === ")

	for {
		var input1 string
		var input2 string

		fmt.Print("Enter Input1 : ")
		fmt.Scanln(&input1)

		var operator string
		fmt.Print("select operation (+,-,*,/, exit, Help): ")
		fmt.Scanln(&operator)

		fmt.Print("Enter Input2 : ")
		fmt.Scanln(&input2)

		command1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Enter Numbers Only")
			continue
		}

		if operator == "Help" {
			fmt.Println("add <a> <b>   → addition")
			fmt.Println("sub <a> <b>   → subtraction")
			fmt.Println("mul <a> <b>   → multiplication")
			fmt.Println("div <a> <b>   → division")
		}

		if operator == "exit" {
			fmt.Println("exited")
			break
		}

		if operator != "+" && operator != "-" && operator != "/" && operator != "*" {
			fmt.Println("Try using Help")
		}

		command2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Enter Numbers Only")
		}

		if operator == "+" {
			fmt.Println("Answer = ", command1+command2)
		} else if operator == "-" {
			fmt.Println("Answer = ", command1-command2)
		} else if operator == "*" {
			fmt.Println("Answer = ", command1*command2)
		} else if operator == "/" {
			if command2 == 0 {
				fmt.Println("Error can't divide by zero")
			} else {
				fmt.Println("Answer = ", command1/command2)
			}
		}
	}
}

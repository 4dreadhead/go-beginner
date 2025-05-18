package main

import (
	"fmt"
)

func main() {
	var transactions []float64
	for {
		var userInput float64
		fmt.Print("Enter transaction: ")
		val, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Println("Error parsing user input:", err)
			fmt.Println(val)
			continue
		}
		if userInput == 0 {
			break
		}
		transactions = append(transactions, userInput)
	}
	summ := 0.0
	for _, value := range transactions {
		summ += value
	}
	fmt.Println("Your transactions:", transactions)
	fmt.Println("Summary transactions:", summ)
}

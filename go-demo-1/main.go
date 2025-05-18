package main

import (
	"fmt"
	"math"
	"strings"
)

const SQUARE = 2

func main() {
	for {
		userHeight, userWeightKg := InputResult()
		bodyMassIndex := CalculateResult(userHeight, userWeightKg)
		OutputResult(bodyMassIndex)
		if NeedToContinue() {
			continue
		} else {
			break
		}
	}
}

func InputResult() (float64, float64) {
	var userHeight, userWeightKg float64
	fmt.Print("Enter height: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter weight: ")
	fmt.Scan(&userWeightKg)

	return userHeight, userWeightKg
}

func CalculateResult(height float64, weight float64) float64 {
	return weight / math.Pow(height, SQUARE)
}

func OutputResult(value float64) {
	switch {
	case value < 16:
		fmt.Println("Very Skin bitch")
	case value < 18.5:
		fmt.Println("Skin bitch")
	case value < 25:
		fmt.Println("Good boy")
	case value < 30:
		fmt.Println("Big boy")
	default:
		fmt.Println("Damn you are a Bus")
	}
	fmt.Printf("Your BMI: %v\n", value)
}

func NeedToContinue() bool {
	var val string
	fmt.Printf("Do you want to continue? ")
	fmt.Scan(&val)

	switch strings.ToLower(val) {
	case "next", "true", "yes", "1", "y":
		return true
	default:
		return false
	}
}

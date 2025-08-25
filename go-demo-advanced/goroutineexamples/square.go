package goroutineexamples

import (
	"fmt"
	rand2 "math/rand/v2"
)

const times = 10

func SquareExample() {
	inputChannel := make(chan int, times)
	outputChannel := make(chan int, times)

	go generateNumber(inputChannel)
	go calculateSquare(inputChannel, outputChannel)

	var responses [times]int

	for i := range times {
		responses[i] = <-outputChannel
	}
	fmt.Printf("Squares: %v\n", responses)
}

func generateNumber(channel chan<- int) {
	var elements [times]int
	for i := range times {
		elements[i] = rand2.IntN(100)
	}
	for _, element := range elements {
		channel <- element
	}
}

func calculateSquare(inputChannel <-chan int, outputChannel chan<- int) {
	for res := range inputChannel {
		outputChannel <- res * res
	}
}

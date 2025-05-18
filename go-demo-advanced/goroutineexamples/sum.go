package goroutineexamples

import "fmt"

func SumExample() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	num := 3
	size := len(arr) / num
	pipe := make(chan int, num)
	sum := 0

	for index := range num {
		go sumPart(arr[index*size:(index+1)*size], pipe)
	}
	for range num {
		sum += <-pipe
	}
	fmt.Printf("Final sum: %v\n", sum)
}

func sumPart(arr []int, pipe chan<- int) {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	pipe <- sum
}

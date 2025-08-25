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

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for i := range 10 {
			chan1 <- i
		}
		close(chan1)
	}()
	go func() {
		for i := range 20 {
			chan2 <- i
		}
		close(chan2)
	}()

	channel1Closed := false
	channel2Closed := false

	for {
		select {
		case res, ok := <-chan1:
			if ok {
				fmt.Println(res)
			} else {
				channel1Closed = true
			}
		case res, ok := <-chan2:
			if ok {
				fmt.Println(res)
			} else {
				channel2Closed = true
			}
		}

		if channel1Closed && channel2Closed {
			break
		}
	}
}

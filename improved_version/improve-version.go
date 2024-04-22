package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var result = []int{}
	channels := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		randon1 := rand.Intn(1000-1) - 1
		randon2 := rand.Intn(1000-1) - 1
		go execute(channels[i], int(randon1), int(randon2))
	}

	for i := 0; i < 10; i++ {
		result = append(result, <-channels[i])
	}

	fmt.Println(result)
}

func execute(ch chan int, value1 int, value2 int) {
	go func() {
		sum := value1 + value2
		time.Sleep(1 * time.Second)
		ch <- sum
	}()
}

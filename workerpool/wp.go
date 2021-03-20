package main

import (
	"fmt"
	"time"
)

func main() {
	numWorkers := 50
	c := make(chan string)

	for i := 0; i < numWorkers; i++ {
		go func(count int) {
			time.Sleep(time.Millisecond * 5000)
			c <- fmt.Sprintf("Worker #%d", count)
		}(i)
	}

	for res := range <-c {
		fmt.Println(res)
	}

}

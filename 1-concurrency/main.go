package main

import (
	"fmt"
	"math/rand"
)

func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			num := rand.Intn(100)
			ch1 <- num
		}
	}()

	go func() {
		defer close(ch2)
		for res := range ch1 {
			ch2 <- res * res
		}
	}()

	for res := range ch2 {
		fmt.Println(res)
	}

}

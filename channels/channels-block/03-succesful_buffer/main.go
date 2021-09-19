package main

import "fmt"

func main() {
	c := make(chan int, 2)

	go func() {
		c <- 1
		c <- 32
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
}

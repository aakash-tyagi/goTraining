package main

import "fmt"

func main() {

	c := make(chan int)
	cr := make(chan<- int) //recive
	cs := make(<-chan int) //send

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// c = cr   specific to general do not assign
	// c = cs
}

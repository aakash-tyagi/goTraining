package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p *person) speak() {
	fmt.Println("Hello I am ", p.Name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		Name: "Aakash",
		Age:  27,
	}

	saySomething(&p1) //we can pass only pointer here
}

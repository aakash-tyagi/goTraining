package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("foo")
	wg.Done()
}

func bar() {
	fmt.Println("bar")
	wg.Done()
}

func main() {
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("goRoutine", runtime.NumGoroutine())

	wg.Add(2)

	go foo()
	go bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("goRoutine", runtime.NumGoroutine())
	wg.Wait()

}

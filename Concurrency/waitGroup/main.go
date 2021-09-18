package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:\t", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:\t", i)
	}
}

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("goRoutines\t", runtime.NumGoroutine())

	go foo()
	bar()

	wg.Add(1)
	wg.Wait()
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("goRoutines\t", runtime.NumGoroutine())

}

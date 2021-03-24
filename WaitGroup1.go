package main

import (
	"fmt"
	"sync"
)

func myfunc(wg *sync.WaitGroup) {
	fmt.Println("Inside my Goroutine")
	wg.Done()
}
func main() {
	fmt.Println("Hello World")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myfunc(&waitgroup)
	waitgroup.Wait()
	fmt.Println("Finished Execution")
}

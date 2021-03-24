package main

import "fmt"

func main() {
	fmt.Println("Anonymouse Goroutine:")
	go func() {
		fmt.Println("MY anonymouse function:")
	}()
	fmt.Scanln()
}

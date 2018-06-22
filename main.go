package main

import "fmt"

func main() {
	c := make(chan string)
	go boring("message", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("received message: %q\n", <-c)
	}
	fmt.Println("completed")
}

package main

import "fmt"

func main() {
	generatorExec()
}

func boringExec() {
	c := make(chan string)
	go boring("message", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("received message: %q\n", <-c)
	}
	fmt.Println("completed")
}

func generatorExec() {
	c := generator("generator")
	for i := 0; i < 5; i++ {
		fmt.Printf("Generator said: %q\n", <-c)
	}
	fmt.Println("Leaving...")
}

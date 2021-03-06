package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timedSearch()
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
	foo := generator("foo")
	bar := generator("bar")
	for i := 0; i < 5; i++ {
		fmt.Println(<-foo)
		fmt.Println(<-bar)
	}
	fmt.Println("Leaving...")
}

func multiplexorExec() {
	c := selectMultiplexor(generator("foo"), generator("bar"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Leaving...")
}

func timedSearch() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := SearchMe("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

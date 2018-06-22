package main

import (
	"fmt"
	"time"
)

func multiplexor(inputA, inputB <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-inputA
		}
	}()
	go func() {
		for {
			c <- <-inputB
		}
	}()
	return c
}

func selectMultiplexor(inputA, inputB <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		// global timeout for the entire loop:
		// timeout := time.After(5 * time.Second)
		for {
			select {
			case s := <-inputA:
				c <- s
			case s := <-inputB:
				c <- s
			case <-time.After(1 * time.Second):
				fmt.Println("Time out")
				return
			}
		}
	}()
	return c
}

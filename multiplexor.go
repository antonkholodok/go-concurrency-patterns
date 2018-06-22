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

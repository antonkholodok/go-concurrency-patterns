package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WaitingMessage struct {
	msg  string
	wait chan bool
}

func generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func waitingGenerator(msg string) <-chan WaitingMessage {
	c := make(chan WaitingMessage)
	go func() {
		waitForIt := make(chan bool)
		for i := 0; ; i++ {
			c <- WaitingMessage{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

package main

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

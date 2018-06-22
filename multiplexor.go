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

func selectMultiplexor(inputA, inputB <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-inputA:
				c <- s
			case s := <-inputB:
				c <- s
			}
		}
	}()
	return c
}

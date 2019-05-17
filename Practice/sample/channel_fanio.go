package main

func main() {
	d := make(chan string, 1)
	h := make(chan string, 1)
	go func() {
		d <- "hello"
	}()
	go func() {
		h <- "world"
	}()
	print(<-fanIn(d, h))

}

// fan in/out  input params only for receive
func fanIn(input1, input2 <-chan string) (<-chan string) {
	c := make(chan string)
	go func() {
		select {
		case s := <-input1:
			c <- s
		case s := <-input2:
			c <- s
		}
	}()
	return c
}

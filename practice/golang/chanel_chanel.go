package main

import "fmt"

func main() {
	g := make(chan int)
	quit := make(chan chan bool)
	go B(g, quit)
	for i:= 0; i < 6; i++{
		g <- i
	}
	wait := make(chan bool)
	quit <- wait
	<-wait
	fmt.Println("main quit")
}
func B(g chan int, quit chan chan bool)  {
	for {
		select {
		case i:= <-g:
			fmt.Println(i + 1)
		case c:= <-quit:
			c <- true
			fmt.Println("B quit")
			return
		}
	}
}

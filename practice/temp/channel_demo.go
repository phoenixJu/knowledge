package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int, 1)
	var a string
	go func() {
		a = "xx"
		<-c1
		<-c2
		fmt.Println(a)
		<-c1
	}()
	c1 <- 1
	fmt.Println(a)
	a = "yy"
	c2 <- 1
	c1 <- 1
}

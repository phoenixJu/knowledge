package main

import (
	"fmt"
	"time"
)

func doService() {
	fmt.Println("I am do main things.")
	time.Sleep(time.Duration(3 * time.Second))
	fmt.Println("main things finished.")
}

func aSyncDoService() chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		fmt.Println("do other things, maybe a long time")
		time.Sleep(time.Duration(6 * time.Second))
		fmt.Println("async finished.")
		c <- struct{}{}
	}()
	return c
}

func main() {
	c := aSyncDoService()
	doService()
	select {
	case <-c:
		fmt.Println("other also finished.")
	case <-time.After(2 * time.Second):
		fmt.Println("timeout error.")
	}
	fmt.Println("exit")
}

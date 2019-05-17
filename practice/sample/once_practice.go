package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceFun := func() {
		fmt.Println("only once")
	}
	done := make(chan bool)
	for i := 0; i < 20; i++ {
		go func() {
			once.Do(onceFun)
			done <- true
		}()
	}
	for i := 0; i < 20; i++ {
		<-done
	}
}

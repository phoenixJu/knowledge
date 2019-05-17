package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for t := range ticker.C {
			fmt.Println("2 second pass... %v", t)
		}
	}()
	time.Sleep(6 * time.Second)
	ticker.Stop()
	time.Sleep(5 * time.Second)
}

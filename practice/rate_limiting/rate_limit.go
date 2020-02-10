package main

import (
	"fmt"
	"time"
)

func Caller() {
	fmt.Println("method invoked !")
}

func main() {
	//d := time.Minute / 100
	//for {
	//	<- time.Tick(d)
	//	go Caller()
	//}
	// 每分钟不超过100次
	d2 := time.Minute  / 10
	ticker := time.NewTicker(d2)
	defer ticker.Stop()
	throttle := make(chan time.Time, 100)
	go func() {
		for t := range ticker.C {
			//throttle = append(throttle, t)
			select {
			case throttle <- t:
			default:
			}
		}
	}()
	for {
		<- throttle
		go Caller()
	}
}

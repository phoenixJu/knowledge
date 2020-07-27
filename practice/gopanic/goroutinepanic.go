package main

import (
	"fmt"
	"time"
)

func main3()  {
	go callerA()
	fmt.Println("finish")
	time.Sleep(8 * time.Second)
	fmt.Println("8 seconds finish")
}

func callerA(){
	fmt.Println("callerA begin")
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("CalcuTrackISRC")
		panic("panic")
	}()
	fmt.Println("callerA end ")
	return
}
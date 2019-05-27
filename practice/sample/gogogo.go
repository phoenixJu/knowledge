package main

import (
	"fmt"
	"time"
)

func forkGo(i int)  {
	fmt.Println(i)
	i++
	a := i
	go func() {
		fmt.Println(a)
		forkGo(a)
	}()
	time.Sleep(6 * time.Second)
}

func main()  {
	forkGo(0)
	time.Sleep(100 * time.Second)
}

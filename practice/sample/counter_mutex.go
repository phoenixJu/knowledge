package main

import (
	"sync"
	"time"
)

func main()  {
	counter := 0
	var mut sync.Mutex
	for i:= 0; i< 5000; i++{
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(3 * time.Duration(1 * time.Second))
	println(counter)
}

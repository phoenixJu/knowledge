package main

import "fmt"

func (it Int64s)Iterator() <-chan int64 {
	ch := make(chan int64, 0)
	go func() {
		defer close(ch)
		for _, i := range it{
			ch <- i
		}
	}()
	return ch
}

func main()  {
	ints := Int64s{1, 2, 5, 8, 13}
	it := ints.Iterator()
	for i:= range it{
		fmt.Println(i)
	}
}

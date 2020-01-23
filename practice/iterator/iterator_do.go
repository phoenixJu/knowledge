package main

import "fmt"

type Int64s []int64

func (i Int64s) Do(fn func(int64))  {
	for _, v := range i{
		fn(v)
	}
}

func main2()  {
	ints := Int64s{1, 2, 5, 8, 13}
	ints.Do(func(i int64) {
		fmt.Println(i)
	})
}
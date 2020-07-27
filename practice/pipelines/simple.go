package main

import (
	"fmt"
	"sync"
	"time"
)

// 如果不起goroutine调用gen，调用者会因为ch而阻塞, 如果不关闭ch，接受者会阻塞
func gen(data ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, d := range data {
			ch <- d
		}
		close(ch)
	}()
	return ch
}

func sq(data <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for d := range data {
			out <- d * d
		}
		close(out)
	}()
	return out
}

func merge(done chan struct{}, chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		defer wg.Done()
		for d := range c {
			select {
			case out <- d:
			case <-done:
				fmt.Println(d)
				return
			}
		}
	}
	wg.Add(len(chs))
	for _, cs := range chs {
		go output(cs) // concurrent
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func mainx() {
//	done := make(chan struct{}, 2)
	done := make(chan struct{})
	defer close(done)
	in := gen(3, 3, 4, 4, 4, 8, 8, 8, 8, 100, 545, 456346, )
	ch2 := sq(in)
	ch3 := sq(in)
	out := merge(done, ch2, ch3)
	fmt.Println(<-out)
	time.Sleep(300000)
}

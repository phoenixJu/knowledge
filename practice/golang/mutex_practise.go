package main

import (
	"fmt"
	"sync"
	"time"
)
// 主协程和其他协程都是独立互不影响地执行，先后顺序不能保证，通常主协程了各个协程之间是通过存放一个空的struct消息的channel来完成同步的
func main() {
	var l sync.Mutex
	c := make(chan struct{}, 2)
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1: 我锁定了2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了，你可以抢了")
		c <- struct{}{}
	}()
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 我锁定了3s")
		time.Sleep(time.Second * 3)
		fmt.Println("goroutine2: 我解锁了，大家可以抢了")
		c <- struct{}{}
	}()
	for i := 0; i < 2; i++ {
		<-c
	}
}

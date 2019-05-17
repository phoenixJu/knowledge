package main

import (
	"fmt"
	"sync"
	"time"
)

// 和waitgroup区别，前者是等大家齐了一起干，cont是预备开始，有点类似监听者

var ct int = 4

func main() {
	ch := make(chan struct{}, 5)
	//新建cond
	var l sync.Mutex
	cond := sync.NewCond(&l)
	for i := 0; i < 5; i++ {
		go func(i int) {
			//争抢锁定
			cond.L.Lock()
			defer func() {
				cond.L.Unlock()
				ch <- struct{}{}
			}()
			//条件是否达成
			for ct > i {
				//这是个生产者只是初始队列为满的
				cond.Wait()
				fmt.Printf("收到一个通知goroutine%d\n", i)
			}
			fmt.Printf("goroutine%d 执行结束\n", i)
		}(i)
	}

	time.Sleep(time.Second * 10)
	fmt.Println("broadcast ...")
	cond.L.Lock()
	ct -= 1
	cond.Broadcast()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.L.Lock()
	ct -= 2
	cond.Signal()
	cond.L.Unlock()
	time.Sleep(time.Second)
	fmt.Println("broadcast ...")
	cond.L.Lock()
	ct -= 1
	cond.Broadcast()
	cond.L.Unlock()

	for i := 0; i < 5; i++ {
		<-ch
	}

}

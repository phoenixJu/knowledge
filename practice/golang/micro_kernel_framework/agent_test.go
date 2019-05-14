package micro_kernel_framework

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type MyCollector struct {
	evtReceiver EventReceiver
	agtCtx context.Context
	stopChan chan struct{}
	name string
	content string
}

func (c *MyCollector) Init(eventReceiver EventReceiver) error {
	fmt.Println("初始化 collector!", c.name)
	c.evtReceiver = eventReceiver
	return nil
}

func (c *MyCollector) Start(agtCtx context.Context) error {
	fmt.Println("启动 collector!", c.name)
	for {
		select {
		// 通过context实现停止功能
		case <- agtCtx.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50)
			// 完成启动后通知receiver
			c.evtReceiver.OnEvent(Event{c.name, c.content})
		}
	}
}

func (c *MyCollector) Stop() error {
	fmt.Println("停止 collector!", c.name)
	select {
	case <- c.stopChan:
		return nil
		case <- time.After(time.Second * 1):
			return errors.New("停止失败，超时了！")
	}
}


func (c *MyCollector) Destroy() error {
	fmt.Println(c.name, "释放资源！")
	return nil
}


func NewCollect(name string, content string)*MyCollector{
	return &MyCollector{
		stopChan: make(chan struct{}),
		name:name,
		content:content,
	}
}

func TestAgent(t *testing.T) {
	agt := NewAgent(100)
	c1 := NewCollect("bytedance", "字节跳动")
	c2 := NewCollect("baidu", "baidu")
	agt.RegisterCollector("字节跳动收集器",c1)
	agt.RegisterCollector("百度收集器", c2)
	agt.Start()
	fmt.Println(agt.Start())
	time.Sleep(time.Second * 1)
	agt.Stop()
	agt.Destroy()
}

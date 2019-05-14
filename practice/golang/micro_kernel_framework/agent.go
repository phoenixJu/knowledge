package micro_kernel_framework

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("当前状态无法完成操作")

type CollectorsError struct {
	CollectorErrors [] error
}

func (ce CollectorsError) Error() string {
	var strs [] string
	for _, err := range ce.CollectorErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ";")
}

type Event struct {
	Source  string
	Content string
}
type EventReceiver interface {
	// 每个collect可以回传
	OnEvent(evt Event)
}
type Collector interface {
	Init(eventReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	state      int
	evtBuf     chan Event
	cancel     context.CancelFunc
	ctx        context.Context
}

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}
	return &agt
}
func (agt *Agent) EventProcessGroutine() {
	var evtSeg [10] Event
	for {
		//连续接收10个信号就会打印一行消息
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuf:
			case <-agt.ctx.Done():
				return
			}
		}
		fmt.Println(evtSeg)
	}
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	// 注册一个collector到agent
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.collectors[name] = collector
	return collector.Init(agt)
}

func (agt *Agent) startCollectors() error {
	var err error
	var errs CollectorsError
	var mutex sync.Mutex
	for name, collector := range agt.collectors {
		// 开启很多goroutine来启动Collector
		go func(name string, collector Collector, ctx context.Context) {
			defer func() { mutex.Unlock() }()
			err = collector.Start(ctx)
			mutex.Lock()
			if err != nil {
				// mutex用来同步CollectorErrors，多个协程共享mutex互斥锁
				errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
			}
		}(name, collector, agt.ctx)
	}
	return errs
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	return errs
}

func (agt *Agent) destryCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destroy(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	// 单独开一个协程监控消息事件
	go agt.EventProcessGroutine()
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStateError
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent) Destroy() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	return agt.destryCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}

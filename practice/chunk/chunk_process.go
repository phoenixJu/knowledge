package chunk

import (
	"math"
	"sync"
)

// chan 本质是支持size=0的阻塞队列
func ChunkProcessor(size int, processure func(start, end int) (interface{}, error),
	reduce func(start, end int, partErr error, partResult interface{}), maxConcurrency int, chunkSize int) {
	if size <= 0 || chunkSize <= 0 || maxConcurrency <= 0 {
		return
	}
	var wg sync.WaitGroup
	var mutx sync.Mutex
	var chs = make(chan struct{}, maxConcurrency)
	chanLimit := int(math.Ceil(float64(size) / float64(chunkSize)))
	wg.Add(chanLimit)
	startIndex := 0
	for {
		endIndex := startIndex + chunkSize
		if endIndex >= size {
			endIndex = size
		}
		chs <- struct{}{}
		go func(s, e int) {
			defer func() {
				wg.Done()
				<- chs
			}()
			ret, err := processure(s, e)
			mutx.Lock()
			reduce(startIndex, endIndex, err, ret)
			mutx.Unlock()
		}(startIndex, endIndex)

		startIndex = startIndex + chunkSize
		if startIndex >= size {
			break
		}
	}
	wg.Wait()
}

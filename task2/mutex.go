package task2

import (
	"sync"
	"sync/atomic"
)

type Count struct {
	sync.Mutex
	counter int
}

func NewCount() *Count {
	return &Count{}
}

func (c *Count) Increase() {
	c.Lock()
	defer c.Unlock()
	c.counter++
}

func (c *Count) value() int {
	return c.counter
}

func TestCount() {
	count := NewCount()
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				count.Increase()
			}
		}(wg)
	}
	wg.Wait()
	println(count.value())
}

func TestAtomic() {
	var count int64
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&count, 1)
			}
		}(wg)
	}
	wg.Wait()
	println(count)
}

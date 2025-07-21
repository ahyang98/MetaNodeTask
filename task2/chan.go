package task2

import (
	"fmt"
	"sync"
)

func producer(ch chan int, wg *sync.WaitGroup, size int) {
	defer close(ch)
	defer wg.Done()
	for i := 1; i <= size; i++ {
		ch <- i
	}

}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-ch:
			if ok {
				fmt.Println(num)
			} else {
				return
			}
		}
	}
}

func TestCommunication() {
	fmt.Println("start")
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go producer(ch, wg, 10)
	go consumer(ch, wg)
	wg.Wait()
	fmt.Println("...")
	wg.Add(2)
	ch = make(chan int, 10)
	go producer(ch, wg, 100)
	go consumer(ch, wg)
	wg.Wait()
	fmt.Println("finish")
}

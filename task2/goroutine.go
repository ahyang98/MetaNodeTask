package task2

import (
	"fmt"
	"sync"
	"time"
)

func PrintNum(isOdd bool, top int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println()
	for i := 1; i <= top; i++ {
		if i%2 == 1 {
			if isOdd {
				fmt.Printf("%d ", i)
			}
		} else {
			if !isOdd {
				fmt.Printf("%d ", i)
			}
		}
	}
	fmt.Println()
}

func TestGoroutine() {
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintNum(true, 10, &wg)
	go PrintNum(false, 10, &wg)
	wg.Wait()
}

type WorkerType func(wg *sync.WaitGroup)

func Schedule(workers []WorkerType) {
	var wg sync.WaitGroup
	for _, worker := range workers {
		wg.Add(1)
		go calDuration(worker, &wg)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Wait()
	fmt.Println("\nfinish")
}

func calDuration(worker WorkerType, wg *sync.WaitGroup) {
	start := time.Now().UnixMicro()
	worker(wg)
	end := time.Now().UnixMicro()
	fmt.Printf("time cost %d ms\n", end-start)
}

func TestSchedule() {
	worker1 := func(wg *sync.WaitGroup) {
		PrintNum(true, 10, wg)
	}
	worker2 := func(wg *sync.WaitGroup) {
		PrintNum(false, 10, wg)
	}
	Schedule([]WorkerType{worker1, worker2})
}

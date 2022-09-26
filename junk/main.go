package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex sync.RWMutex
	wg    sync.WaitGroup
)

func main() {
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		mutex.RLock()
		defer wg.Done()
		defer mutex.RUnlock()

		fmt.Scanln()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Millisecond)
		wg.Add(1)
		mutex.RLock()
		defer wg.Done()
		defer mutex.RUnlock()

		fmt.Printf("waiting for input...")
	}(&wg)

	time.Sleep(time.Millisecond)
	wg.Wait()
}

package main

import (
	"fmt"
	"time"
)

type RWMutex struct {
	ch      chan struct{}
	rch     chan struct{}
	counter int
}

func (rw *RWMutex) Lock() {
	rw.rch <- struct{}{}
	rw.ch <- struct{}{}
	fmt.Print("Lock\n")
}

func (rw *RWMutex) Unlock() {
	fmt.Print("Unlock\n")
	<-rw.rch
	<-rw.ch
}

func (rw *RWMutex) RLock() {
	rw.counter++
	if rw.counter == 1 {
		rw.rch <- struct{}{}
	}
	rw.ch <- struct{}{}
	<-rw.ch
	fmt.Print("Rlock\n")
}

func (rw *RWMutex) RUnlock() {
	fmt.Print("RUnlock\n")
	rw.counter--
	if rw.counter == 0 {
		<-rw.rch
	}
}

func New() *RWMutex {
	return &RWMutex{
		ch:  make(chan struct{}, 1),
		rch: make(chan struct{}, 1),
	}
}

func main() {
	mtx := New()
	sync := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(mutex *RWMutex, sync chan bool, number int) {
			mtx.Lock()
			//fmt.Printf("%d: Im writing... | number of readers: %d\n", number, mtx.counter)
			time.Sleep(time.Second)
			mtx.Unlock()
			sync <- true
		}(mtx, sync, i)

		go func(mutex *RWMutex, sync chan bool, number int) {
			time.Sleep(3 * time.Second)
			mtx.RLock()
			//fmt.Printf("%d: Im reading... | number of readers: %d\n", number, mtx.counter)
			time.Sleep(time.Second)
			mtx.RUnlock()
			sync <- true
		}(mtx, sync, i)
	}
	for i := 0; i < 10; i++ {
		<-sync
	}
	for i := 0; i < 10; i++ {
		<-sync
	}
}

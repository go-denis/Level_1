package main

import (
	"fmt"
	"sync"
)

type ConCount struct {
	mu  sync.Mutex
	cnt int
}

func main() {

	// Счетчик с Mutex
	c := ConCount{cnt: 0}

	// Атомарный счетчик
	var c1 int64

	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				//atomic.AddInt64(&c1, 1)

				c.mu.Lock()
				c.cnt++
				c.mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("atomic", c1)
	fmt.Println("mutex ", c.cnt)
}

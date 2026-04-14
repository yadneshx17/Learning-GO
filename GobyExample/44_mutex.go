package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock() // unlock at end of the functino using defer statement
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}
	var wg sync.WaitGroup

	doIncrement := func(name string, count int) {
		for range count {
			c.inc(name)
		}
	}

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	wg.Wait()
	fmt.Println(c.counters)
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func doPrint(wg *sync.WaitGroup, message string) {
// 	defer wg.Done()
// 	fmt.Println(message)
// }

// A.60 sync mutex

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Add(int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counter) Value() int {
	return c.val
}

func main() {

	// A.59 sync waitgroup

	// runtime.GOMAXPROCS(2)

	// var wg sync.WaitGroup

	// for i := 0; i < 5; i++ {
	// 	var data = fmt.Sprintf("data %d", i)

	// 	wg.Add(1)

	// 	go doPrint(&wg, data)
	// }

	// wg.Wait()

	// A.60 sync mutex

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter
	var mtx sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				mtx.Lock()
				meter.Add(1)
				mtx.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println((meter.Value()))

}

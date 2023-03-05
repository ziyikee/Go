package test

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main2() {
	test2()
}

func test2() {
	// ruleid: waitgroup-add-called-inside-goroutine
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	var x int32 = 0
	wg1.Add(1)
	for i := 0; i < 100; i++ {
		go func() {
			wg1.Add(1)
			go func() {
				return
			}()
			atomic.AddInt32(&x, 1)
			wg1.Done()
			wg2.Add(1)
		}()
	}

	fmt.Println("Wait ...")
	wg1.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}

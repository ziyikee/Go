package test

import (
	"fmt"
	"sync"
)

func main2() {
	test2()
}

func test2() {
	// ruleid: waitgroup-add-called-inside-goroutine
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	wg3 := sync.WaitGroup{}
	wg3.Wait()
	wg1.Add(1)
	for i := 0; i < 100; i++ {
		go func() {
			wg1.Add(1)
			go func() {
				wg1.Add(1)
				go func() {
					return
				}()
				return
			}()
			wg1.Done()
			wg2.Add(1)
			addCall(wg2)
		}()
	}

	fmt.Println("Wait ...")
	wg1.Wait()
}

func addCall(wg2 sync.WaitGroup) {
	wg2.Add(1)
}

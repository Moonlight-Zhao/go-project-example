package concurrence

import (
	"fmt"
	"sync"
)

func hello(i int) {
	println("hello world : " + fmt.Sprint(i))
}

func ManyGo() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}

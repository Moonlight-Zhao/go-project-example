package attention

import (
	"time"
)

func closure() {
	for i := 0; i < 3; i++ {
		go func() {
			println(i)
		}()
	}
	time.Sleep(3 * time.Second)
}

func closure1() {
	for i := 0; i < 3; i++ {
		go func(j int) {
			println(j)
		}(i)
	}
	time.Sleep(3 * time.Second)
}






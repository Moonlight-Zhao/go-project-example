package benchmark

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"math/rand"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i
	}
}

func Select() int {
	return rand.Intn(10)
}

func FastSelect() int {
	return fastrand.Intn(10)
}
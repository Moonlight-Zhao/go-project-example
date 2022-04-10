package attention

import "fmt"

func AppendInt() {
	intArray := [3]int64{1, 2, 3}
	func(arr [3]int64) {
		arr[2] = 4
		fmt.Println("inner func array:",arr)
	}(intArray)

	fmt.Println("outer func array:",intArray)
}







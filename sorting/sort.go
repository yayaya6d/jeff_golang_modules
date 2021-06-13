package sorting

import "reflect"

func Sort(data interface{}, less func(i, j int) bool) {
	rv := reflect.ValueOf(data)
	swap := reflect.Swapper(data)
	length := rv.Len()
	insertionSort_func(lessSwap{less, swap}, 0, length)
}

type lessSwap struct {
	Less func(i, j int) bool
	Swap func(i, j int)
}

func insertionSort_func(data lessSwap, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

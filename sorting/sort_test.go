package sorting

import (
	"testing"
)

func Test_Sort_SortIntSlice(t *testing.T) {
	arr := []int{8, 2, 6, 3, 1, 4}
	Sort(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	expected := []int{1, 2, 3, 4, 6, 8}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("unexpected value in index %d, expect %d, got %d", i, expected[i], arr[i])
		}
	}
}

type complex struct {
	i int
	s string
	b bool
}

func Test_Sort_SortStructSlice(t *testing.T) {
	arr := []complex{{1, "", true}, {0, "", true}, {3, "3", true}, {2, "", false}, {2, "", true}}

	Sort(arr, func(i, j int) bool {
		if arr[i].i < arr[j].i {
			return true
		}

		return arr[i].b && !arr[j].b
	})

	expected := []complex{{0, "", true}, {1, "", true}, {2, "", true}, {2, "", false}, {3, "3", true}}

	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("unexpected value in index %d.", i)
		}
	}
}

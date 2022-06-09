package slices

import (
	"github.com/matryer/is"
	"testing"
)

type tester struct {
	sliceOne interface{}
	sliceTwo interface{}
	want     interface{}
}

func TestConcat(t *testing.T) {
	is := is.New(t)
	f := func(given [][]int, expected []int) {
		is.Equal(Concat(given...), expected)
	}
	f([][]int{}, []int{})
	f([][]int{{}}, []int{})
	f([][]int{{1}}, []int{1})
	f([][]int{{1}, {}}, []int{1})
	f([][]int{{}, {1}}, []int{1})
	f([][]int{{1, 2}, {3, 4, 5}}, []int{1, 2, 3, 4, 5})
}

func TestZip(t *testing.T) {
	is := is.New(t)
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range Zip(given...) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		is.Equal(expected, actual)
	}
	f([][]int{}, [][]int{})
	f([][]int{{1}, {2}}, [][]int{{1, 2}})
	f([][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}})
	f([][]int{{1, 2, 3}, {4, 5}}, [][]int{{1, 4}, {2, 5}})
}

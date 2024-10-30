package main

import (
	"reflect"
	"testing"
)

func TestDs(t *testing.T){
	t.Run("test sum array", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("return number of slices", func(t *testing.T){
		got := SumAll([]int{1, 2, 3, 4}, []int{1, 2})
		want := []int{10,3}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v  wanr %v", got, want)
		}
	})

	t.Run("sum all tails", func(t *testing.T){
		got := SumAllTails([]int{1, 2}, []int{1, 2, 3})
		want := []int{2, 5}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely handle empty slices", func(t *testing.T){
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}